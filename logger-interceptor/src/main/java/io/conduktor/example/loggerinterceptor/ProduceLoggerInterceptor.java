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
import lombok.extern.slf4j.Slf4j;
import org.apache.kafka.common.requests.ProduceRequest;

import java.util.concurrent.CompletableFuture;
import java.util.concurrent.CompletionStage;

@Slf4j
public class ProduceLoggerInterceptor implements Interceptor<ProduceRequest> {
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
                log.warn("Record to partition {}  of topic {} is {}", partitionData.index(), data.name(), RecordUtils.readRecords(partitionData.records()));
            });
        });

        return CompletableFuture.completedFuture(input);
    }

    private String getAuditEntry(InterceptorContext interceptorContext) {
        return interceptorContext.requestHeader().clientId();
    }
}