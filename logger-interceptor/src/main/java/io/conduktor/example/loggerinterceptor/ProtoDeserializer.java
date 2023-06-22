package io.conduktor.example.loggerinterceptor;

import com.google.protobuf.Descriptors;
import com.google.protobuf.DynamicMessage;
import com.squareup.wire.schema.internal.parser.*;
import io.confluent.kafka.schemaregistry.client.SchemaRegistryClient;
import io.confluent.kafka.schemaregistry.client.rest.exceptions.RestClientException;
import io.confluent.kafka.schemaregistry.protobuf.MessageIndexes;
import io.confluent.kafka.schemaregistry.protobuf.ProtobufSchema;
import io.confluent.kafka.serializers.protobuf.KafkaProtobufDeserializer;
import io.confluent.kafka.serializers.protobuf.ProtobufSchemaAndValue;
import org.apache.kafka.common.config.ConfigException;
import org.apache.kafka.common.errors.InvalidConfigurationException;
import org.apache.kafka.common.errors.SerializationException;

import java.io.ByteArrayInputStream;
import java.io.IOException;
import java.nio.ByteBuffer;
import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;

public class ProtoDeserializer extends KafkaProtobufDeserializer<DynamicMessage> {
    public ProtoDeserializer(SchemaRegistryClient client) {
        super(client);
    }

    @Override
    protected Object deserialize(boolean includeSchemaAndVersion, String topic, Boolean isKey, byte[] data) throws SerializationException, InvalidConfigurationException {
        if (this.schemaRegistry == null) {
            throw new InvalidConfigurationException("SchemaRegistryClient not found. You need to configure the deserializer or use deserializer constructor with SchemaRegistryClient.");
        } else if (data == null) {
            return null;
        } else {
            ByteBuffer buffer = this.getByteBuffer(data);
            int id = buffer.getInt();

            try {
                MessageIndexes indexes = MessageIndexes.readFrom(buffer);
                ProtobufSchema schema = (ProtobufSchema) this.schemaRegistry.getSchemaById(id);
                String name = schema.toMessageName(indexes);
                schema = schema.copy(name);

                int length = buffer.limit() - 1 - 4;
                int start = buffer.position() + buffer.arrayOffset();
                Object value;
                if (this.parseMethod != null) {
                    try {
                        value = this.parseMethod.invoke(null, buffer);
                    } catch (Exception var15) {
                        throw new ConfigException("Not a valid protobuf builder", var15);
                    }
                } else {
                    Descriptors.Descriptor descriptor = schema.toDescriptor();
                    if (descriptor == null) {
                        throw new SerializationException("Could not find descriptor with name " + schema.name());
                    }

                    value = DynamicMessage.parseFrom(descriptor, new ByteArrayInputStream(buffer.array(), start, length));
                }

                if (includeSchemaAndVersion) {
                    return new ProtobufSchemaAndValue(schema, value);
                } else {
                    return value;
                }
            } catch (RuntimeException | IOException var16) {
                throw new SerializationException("Error deserializing Protobuf message for id " + id, var16);
            } catch (RestClientException var17) {
                throw toKafkaException(var17, "Error retrieving Protobuf schema for id " + id);
            }
        }
    }

    public Set<TopicRecordField> getRecordFieldValuesAndTags(String topic, byte[] recordData) {
        ProtobufSchemaAndValue schemaAndValue = (ProtobufSchemaAndValue) deserialize(true, topic, this.isKey, recordData);
        DynamicMessage msg = (DynamicMessage) schemaAndValue.getValue();
        Map<String, Object> fields = new HashMap<>();

        msg.getAllFields().forEach((field, object) -> {
            fields.put(field.getName(), msg.getField(field));
        });

        Map<String, Map<String, String>> fieldTags = new HashMap<>();

        ProtoFileElement protoFileElement = schemaAndValue.getSchema().rawSchema();
        String packageNamePrefix = protoFileElement.getPackageName() + ".";
        String msgType = msg.getDescriptorForType().getFullName();

        schemaAndValue.getSchema().rawSchema().getTypes().forEach((TypeElement typeElement) -> {
            MessageElement msgElement = (MessageElement) typeElement;

            // only extract tags for the message type matching the type of the message we parsed
            // not handling complex message fields for now
            if (!msgType.equals(packageNamePrefix + msgElement.getName())) return;

            msgElement.getFields().forEach((FieldElement field) -> {
                Map<String, String> tags = new HashMap<>();

                field.getOptions().forEach((OptionElement optionElement) -> {
                    tags.put(optionElement.getName(), optionElement.getValue().toString());
                });

                fieldTags.put(field.getName(), tags);
            });
        });

        Set<TopicRecordField> res = new HashSet();
        fields.forEach((String name, Object value) -> {
            res.add(new TopicRecordField(name, value, fieldTags.getOrDefault(name, new HashMap<>())));
        });

        return res;
    }
}
