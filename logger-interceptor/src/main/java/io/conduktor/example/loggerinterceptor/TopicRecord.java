package io.conduktor.example.loggerinterceptor;

import io.confluent.kafka.schemaregistry.protobuf.ProtobufSchema;

import java.util.Set;

public class TopicRecord {
    // for simplicity, we assume only string keys if present
    private final String key;
    private final Set<TopicRecordField> fields;

    private final ProtobufSchema schema;

    private final int schemaId;
    private final boolean shouldModify;

    public TopicRecord(String key, Set<TopicRecordField> fields, ProtobufSchema schema, int schemaId) {
        this.key = key;
        this.fields = fields;
        this.schema = schema;
        this.schemaId = schemaId;
        this.shouldModify = shouldModify(fields);
    }

    private static boolean shouldModify(Set<TopicRecordField> fields) {
        for (TopicRecordField field : fields) {
            if (field.isShouldModify()) return true;
        }

        return false;
    }

    public String getKey() {
        return key;
    }

    public Set<TopicRecordField> getFields() {
        return fields;
    }

    public ProtobufSchema getSchema() {
        return schema;
    }

    public int getSchemaId() {
        return schemaId;
    }

    public boolean isShouldModify() {
        return shouldModify;
    }
}
