/*
 * Copyright 2023 Conduktor, Inc
 *
 * Licensed under the Conduktor Community License (the "License"); you may not use
 * this file except in compliance with the License.  You may obtain a copy of the
 * License at
 *
 * https://www.conduktor.io/conduktor-community-license-agreement-v1.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

package io.conduktor.example.loggerinterceptor;

import com.google.protobuf.Descriptors;
import com.google.protobuf.DynamicMessage;
import com.hellofresh.GatewayEncryption;
import io.confluent.kafka.schemaregistry.protobuf.ProtobufSchema;
import org.apache.kafka.common.header.Header;
import org.apache.kafka.common.header.internals.RecordHeader;
import org.apache.kafka.common.record.*;

import java.nio.ByteBuffer;
import java.nio.charset.StandardCharsets;
import java.util.*;
import java.util.concurrent.atomic.AtomicInteger;

public class RecordUtils {

    public static BaseRecords addHeaderToRecords(BaseRecords records, String key, String value) {
        var batchMap = new LinkedHashMap<RecordBatch, List<RecordAndOffset>>();
        AtomicInteger batchesTotalSize = new AtomicInteger();
        ((MemoryRecords) records).batches().forEach(batch -> {
            batchesTotalSize.addAndGet(batch.sizeInBytes());
            var newRecords = new LinkedList<RecordAndOffset>();
            batch.forEach(record -> {

                var newRecord = new SimpleRecord(
                        record.timestamp(),
                        record.key(),
                        record.value(),
                        addHeader(record.headers(), key, value)
                );
                newRecords.add(new RecordAndOffset(newRecord, record.offset()));
            });
            batchMap.put(batch, newRecords);
        });
        var recordsOnly = batchMap.values()
                .stream()
                .flatMap(recordAndOffsets -> recordAndOffsets.stream()
                        .map(RecordAndOffset::record))
                .toList();
        int sizeEstimate = AbstractRecords.estimateSizeInBytes(RecordBatch.CURRENT_MAGIC_VALUE, CompressionType.NONE, recordsOnly);
        var buf = ByteBuffer.allocate(sizeEstimate + batchesTotalSize.get());
        for (var entry : batchMap.entrySet()) {
            writeRecords(buf, entry.getValue(), entry.getKey());
        }
        buf.flip();
        return MemoryRecords.readableRecords(buf);

    }

    /**
     * Processes a batch, unpacks the raw record and find fields that should be encrypted/decrypted
     * @param topic
     * @param records
     * @param deserializer
     * @param serializer
     * @param encryptor
     * @param encrypt When true, relevant fields in records in the batch is encrypted, otherwise, they are decrypted
     * @return
     * @throws Exception
     */
    public static BaseRecords processBatches(String topic, BaseRecords records, ProtoDeserializer deserializer, ProtoSerializer serializer, GatewayEncryption encryptor, boolean encrypt) throws Exception {
        var batchMap = new LinkedHashMap<RecordBatch, List<RecordAndOffset>>();
        AtomicInteger batchesTotalSize = new AtomicInteger();

        for (MutableRecordBatch batch : ((MemoryRecords) records).batches()) {
            batchesTotalSize.addAndGet(batch.sizeInBytes());
            var newRecords = new LinkedList<RecordAndOffset>();

            for (Iterator<org.apache.kafka.common.record.Record> it = batch.iterator(); it.hasNext(); ) {
                org.apache.kafka.common.record.Record record = it.next();

                // unpack record to see if we need to operate on one of its fields
                TopicRecord unpackedRecord = deserializer.getRecord(topic, record);

                ByteBuffer recordKey, recordValue;
                Header[] recordHeaders;

                if (unpackedRecord.isShouldModify()) {
                    ProtobufSchema schema = unpackedRecord.getSchema();
                    Descriptors.Descriptor msgDescriptor = schema.toDescriptor();
                    DynamicMessage.Builder modifiedMessageBuilder = DynamicMessage.newBuilder(msgDescriptor);
                    Map<String, Descriptors.FieldDescriptor> fieldDescriptorMap = new HashMap<>();
                    msgDescriptor.getFields().forEach(field -> {
                        fieldDescriptorMap.put(field.getName(), field);
                    });

                    for (TopicRecordField field : unpackedRecord.getFields()) {
                        modifiedMessageBuilder.setField(
                                fieldDescriptorMap.get(field.getName()),
                                field.isShouldModify() ?
                                        (encrypt ?
                                                encryptor.encryptAES(field.getValue().toString(), unpackedRecord.getKey()) :
                                                tryDecrypt(encryptor, unpackedRecord, field)) :
                                        field.getValue()
                        );
                    }

                    recordValue = ByteBuffer.wrap(serializer.serialize(topic, modifiedMessageBuilder.build(), schema, unpackedRecord.getSchemaId()));
                    recordKey = StandardCharsets.UTF_8.encode(unpackedRecord.getKey());

                    recordHeaders = addHeader(record.headers(), "gwModified", "true");
                } else {
                    recordKey = record.key();
                    recordValue = record.value();
                    recordHeaders = record.headers();
                }

                newRecords.add(new RecordAndOffset(new SimpleRecord(
                        record.timestamp(),
                        recordKey,
                        recordValue,
                        recordHeaders
                ), record.offset()));
            }

            batchMap.put(batch, newRecords);
        }

        var recordsOnly = batchMap.values()
                .stream()
                .flatMap(recordAndOffsets -> recordAndOffsets.stream()
                        .map(RecordAndOffset::record))
                .toList();
        int sizeEstimate = AbstractRecords.estimateSizeInBytes(RecordBatch.CURRENT_MAGIC_VALUE, CompressionType.NONE, recordsOnly);
        var buf = ByteBuffer.allocate(sizeEstimate + batchesTotalSize.get());
        for (var entry : batchMap.entrySet()) {
            writeRecords(buf, entry.getValue(), entry.getKey());
        }

        buf.flip();
        return MemoryRecords.readableRecords(buf);
    }

    private static String tryDecrypt(GatewayEncryption encryptor, TopicRecord unpackedRecord, TopicRecordField field) {
        String plaintext = field.getValue().toString();

        try {
            return encryptor.decryptAES(plaintext, unpackedRecord.getKey());
        } catch (Exception e) {
            return plaintext;
        }
    }

    public static List<TopicRecord> readRecords(String topic, BaseRecords records, ProtoDeserializer deserializer) {
        List<TopicRecord> res = new ArrayList<>();

        ((MemoryRecords) records).batches().forEach(batch -> {
            batch.forEach(record -> {
                res.add(deserializer.getRecord(topic, record));
            });
        });

        return res;
    }

    private static Header[] addHeader(Header[] headers, String key, String value) {
        var listHeaders = new ArrayList<>(List.of(headers));
        listHeaders.add(new RecordHeader(key, value.getBytes()));
        return listHeaders.toArray(new Header[0]);
    }

    private static void writeRecords(ByteBuffer buf, List<RecordAndOffset> recordAndOffsets, RecordBatch batch) {
        long baseOffset = recordAndOffsets.stream()
                .map(RecordAndOffset::offset)
                .min(Comparator.comparingLong(e -> e)).orElse(0L);
        try (var builder = new MemoryRecordsBuilder(
                buf,
                batch.magic(),
                batch.compressionType(),
                batch.timestampType(),
                baseOffset,
                -1L,
                batch.producerId(),
                batch.producerEpoch(),
                batch.baseSequence(),
                batch.isTransactional(),
                batch.isControlBatch(),
                batch.partitionLeaderEpoch(),
                buf.capacity()
        )) {

            for (var recordAndOffset : recordAndOffsets) {
                if (batch.isControlBatch()) {
                    builder.appendControlRecordWithOffset(recordAndOffset.offset(), recordAndOffset.record());
                } else {
                    builder.appendWithOffset(recordAndOffset.offset(), recordAndOffset.record());
                }
            }
            builder.build();
        }
    }

    private record RecordAndOffset(SimpleRecord record, Long offset) {
    }
}
