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

import io.conduktor.gateway.interceptor.Interceptor;
import io.conduktor.gateway.interceptor.InterceptorContext;
import io.confluent.kafka.schemaregistry.client.SchemaRegistryClient;
import lombok.extern.slf4j.Slf4j;
import org.apache.kafka.common.requests.ProduceRequest;

import java.util.List;
import java.util.Set;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.CompletionStage;

@Slf4j
public class ProduceLoggerInterceptor implements Interceptor<ProduceRequest> {
    private final SchemaRegistryClient schemaRegistryClient;

    public ProduceLoggerInterceptor(SchemaRegistryClient schemaRegistryClient) {
        this.schemaRegistryClient = schemaRegistryClient;
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
                List<Set<TopicRecordField>> records = RecordUtils.readRecords(data.name(), partitionData.records(), new ProtoDeserializer(schemaRegistryClient));
                StringBuilder sb = new StringBuilder();

                for (int i = 0; i < records.size(); i++) {
                    sb.append("Record ").append(i).append(": ");

                    records.get(i).forEach((TopicRecordField field) -> {
                        sb.append(field.getName())
                                .append(" = ")
                                .append(field.getValue())
                                .append(" [ ");

                        field.getTags().forEach((String key, String value) -> {
                            sb.append(key)
                                    .append("=")
                                    .append(value)
                                    .append(",");
                        });

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