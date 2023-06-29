package io.conduktor.example.loggerinterceptor;

/**
 * Represents a field in a topic record
 */
public class TopicRecordField {
    private final String name;
    private final Object value;

    private final boolean shouldModify;

    public TopicRecordField(String name, Object value, boolean shouldModify) {
        this.name = name;
        this.value = value;
        this.shouldModify = shouldModify;
    }

    public String getName() {
        return name;
    }

    public Object getValue() {
        return value;
    }

    public boolean isShouldModify() {
        return shouldModify;
    }
}

