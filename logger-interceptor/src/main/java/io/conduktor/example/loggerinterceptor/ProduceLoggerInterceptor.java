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

import com.hellofresh.GatewayEncryption;
import io.conduktor.gateway.interceptor.Interceptor;
import io.conduktor.gateway.interceptor.InterceptorContext;
import io.confluent.kafka.schemaregistry.client.SchemaRegistryClient;
import lombok.extern.slf4j.Slf4j;
import org.apache.kafka.common.requests.ProduceRequest;

import java.util.List;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.CompletionStage;
import java.util.concurrent.atomic.AtomicBoolean;

@Slf4j
public class ProduceLoggerInterceptor implements Interceptor<ProduceRequest> {
    private final SchemaRegistryClient schemaRegistryClient;

    private final GatewayEncryption encryptor;

    public ProduceLoggerInterceptor(SchemaRegistryClient schemaRegistryClient,
                                    GatewayEncryption encryptor) {

        this.schemaRegistryClient = schemaRegistryClient;
        this.encryptor = encryptor;
    }

    @Override
    public CompletionStage<ProduceRequest> intercept(ProduceRequest input, InterceptorContext interceptorContext) {
        var version = interceptorContext.requestHeader().apiVersion();
        log.warn("Produce was called with version: {}", version);

        // add a message header to every message
        // for each topic
        input.data().topicData().forEach(topicProduceData -> {
            // for each partition
            topicProduceData.partitionData().forEach(partitionProduceData -> {
                partitionProduceData.setRecords(RecordUtils.addHeaderToRecords(
                        partitionProduceData.records(),
                        "fromClient", getAuditEntry(interceptorContext)));
            });
        });

        input.data().topicData().forEach(data -> {
            data.partitionData().forEach(partitionData -> {
                List<TopicRecord> records = RecordUtils.readRecords(data.name(), partitionData.records(),
                        new ProtoDeserializer(schemaRegistryClient));
                StringBuilder sb = new StringBuilder();

                for (int i = 0; i < records.size(); i++) {
                    sb.append("Record ").append(i).append(": ");
                    TopicRecord record = records.get(i);

                    record.getFields().forEach((TopicRecordField field) -> {
                        AtomicBoolean shouldEncrypt = new AtomicBoolean(true);
                        StringBuilder tagsSb = new StringBuilder(" [ ");

                        field.getTags().forEach((String key, String value) -> {
                            if (key.equals("hellofresh.tags.data_protection.v1.treatment") && value.equals("TREATMENT_NONE"))
                                shouldEncrypt.set(false);

                            tagsSb.append(key).append("=").append(value).append(",");
                        });

                        try {
                            sb.append(field.getName())
                                    .append(" = ")
                                    .append(shouldEncrypt.get() ? encryptor.encryptAES(field.getValue().toString(), record.getKey()) : field.getValue());
                        } catch (Exception e) {
                            log.error(e.getMessage());
                            throw new RuntimeException(e);
                        }

                        sb.append(tagsSb);
                        sb.append("]; ");
                    });
                }

                log.warn("Record to partition {}  of topic {} is {}",
                        partitionData.index(), data.name(), sb);
            });
        });

        return CompletableFuture.completedFuture(input);
    }

    private String getAuditEntry(InterceptorContext interceptorContext) {
        return interceptorContext.requestHeader().clientId();
    }
}