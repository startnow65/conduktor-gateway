package io.conduktor.example.loggerinterceptor;

import java.util.Map;

/**
 * Represents a field in a topic record
 */
public class TopicRecordField {
    private final String name;
    private final Object value;

    // Tag key/value pairs
    private final Map<String, String> tags;

    public TopicRecordField(String name, Object value, Map<String, String> tags) {
        this.name = name;
        this.value = value;
        this.tags = tags;
    }

    public String getName() {
        return name;
    }

    public Object getValue() {
        return value;
    }

    public Map<String, String> getTags() {
        return tags;
    }
}

