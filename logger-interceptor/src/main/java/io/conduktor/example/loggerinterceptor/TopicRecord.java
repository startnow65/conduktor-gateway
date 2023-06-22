package io.conduktor.example.loggerinterceptor;

import java.util.Set;

public class TopicRecord {
    // for simplicity, we assume only string keys if present
    private final String key;
    private final Set<TopicRecordField> fields;

    public TopicRecord(String key, Set<TopicRecordField> fields) {
        this.key = key;
        this.fields = fields;
    }

    public String getKey() {
        return key;
    }

    public Set<TopicRecordField> getFields() {
        return fields;
    }
}
