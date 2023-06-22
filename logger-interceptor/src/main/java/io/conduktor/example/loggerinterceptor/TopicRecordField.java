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

    private final boolean shouldModify;

    public TopicRecordField(String name, Object value, Map<String, String> tags) {
        this.name = name;
        this.value = value;
        this.tags = tags;
        this.shouldModify = shouldModify(tags);
    }

    private static boolean shouldModify(Map<String, String> tags) {
        for (Map.Entry<String, String> fieldTag : tags.entrySet()) {
            if (fieldTag.getKey().equals("hellofresh.tags.data_protection.v1.treatment")
                    && fieldTag.getValue().equals("TREATMENT_NONE"))
                return false;
        }

        return true;
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

    public boolean isShouldModify() {
        return shouldModify;
    }
}

