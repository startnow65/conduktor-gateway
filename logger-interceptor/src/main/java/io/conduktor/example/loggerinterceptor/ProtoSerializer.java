package io.conduktor.example.loggerinterceptor;

import com.google.protobuf.DynamicMessage;
import com.google.protobuf.Message;
import io.confluent.kafka.schemaregistry.client.SchemaRegistryClient;
import io.confluent.kafka.schemaregistry.client.rest.entities.RuleMode;
import io.confluent.kafka.schemaregistry.protobuf.MessageIndexes;
import io.confluent.kafka.schemaregistry.protobuf.ProtobufSchema;
import io.confluent.kafka.serializers.protobuf.KafkaProtobufSerializer;
import org.apache.kafka.common.errors.InvalidConfigurationException;
import org.apache.kafka.common.errors.SerializationException;
import org.apache.kafka.common.errors.TimeoutException;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.InterruptedIOException;
import java.nio.ByteBuffer;

public class ProtoSerializer extends KafkaProtobufSerializer<DynamicMessage> {
    public ProtoSerializer(SchemaRegistryClient client) {
        super(client);
    }

    public byte[] serialize(String topic, Message record, ProtobufSchema schema, int schemaId) {
        if (schema == null) {
            throw new InvalidConfigurationException("Schema not provided");
        } else if (record == null) {
            return null;
        } else {
            String restClientErrorMsg = "";

            byte[] res;
            try {
                record = (Message) this.executeRules(schema.name(), topic, null, RuleMode.WRITE, null, schema, record);
                ByteArrayOutputStream out = new ByteArrayOutputStream();
                out.write(0);
                out.write(ByteBuffer.allocate(4).putInt(schemaId).array());
                MessageIndexes indexes = schema.toMessageIndexes(record.getDescriptorForType().getFullName(), this.normalizeSchema);
                out.write(indexes.toByteArray());
                record.writeTo(out);
                byte[] bytes = out.toByteArray();
                out.close();
                res = bytes;
            } catch (InterruptedIOException var20) {
                throw new TimeoutException("Error serializing Protobuf message", var20);
            } catch (RuntimeException | IOException var21) {
                throw new SerializationException("Error serializing Protobuf message", var21);
            } finally {
                this.postOp(record);
            }

            return res;
        }
    }
}
