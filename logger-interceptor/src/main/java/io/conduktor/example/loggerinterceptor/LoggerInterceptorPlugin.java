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
import io.conduktor.gateway.interceptor.InterceptorProvider;
import io.conduktor.gateway.interceptor.Plugin;
import io.confluent.kafka.schemaregistry.client.CachedSchemaRegistryClient;
import io.confluent.kafka.schemaregistry.client.SchemaRegistryClient;
import io.confluent.kafka.schemaregistry.protobuf.ProtobufSchemaProvider;
import lombok.extern.slf4j.Slf4j;
import org.apache.kafka.common.requests.*;

import java.util.List;
import java.util.Map;

@Slf4j
public class LoggerInterceptorPlugin implements Plugin {
    static String CONFIG_SCHEMA_REGISTRY_URL = "schemaRegistryUrl";
    static String CONFIG_VAULT_URL = "vaultUrl";
    static String CONFIG_VAULT_TOKEN = "vaultToken";
    static String CONFIG_VAULT_KV_PATH = "vaultKVPath";
    static String CONFIG_FIELD_TAG_KEY_NO_ENCRYPT = "fieldTagKeyNoEncrypt";
    static String CONFIG_FIELD_TAG_VALUE_NO_ENCRYPT = "fieldTagValueNoEncrypt";

    @Override
    public List<InterceptorProvider<?>> getInterceptors(Map<String, Object> config) {
        String prefix = "";
        var loggingStyle = config.get("loggingStyle");
        if (loggingStyle.equals("obiWan")) {
            prefix = "Hello from cloud comms";
        }

        SchemaRegistryClient schemaRegistryClient = new CachedSchemaRegistryClient(
                List.of(config.get(CONFIG_SCHEMA_REGISTRY_URL).toString()), 5000,
                List.of(new ProtobufSchemaProvider()), null
        );

        GatewayEncryption encryptor = new GatewayEncryption(
                config.get(CONFIG_VAULT_URL).toString(),
                config.get(CONFIG_VAULT_TOKEN).toString(),
                config.get(CONFIG_VAULT_KV_PATH).toString());

        ProtoDeserializer deserializer = new ProtoDeserializer(schemaRegistryClient,
                config.get(CONFIG_FIELD_TAG_KEY_NO_ENCRYPT).toString(),
                config.get(CONFIG_FIELD_TAG_VALUE_NO_ENCRYPT).toString());

        ProtoSerializer serializer = new ProtoSerializer(schemaRegistryClient);

        return List.of(
                new InterceptorProvider<>(AbstractRequestResponse.class, new AllLoggerInterceptor(prefix)),
                new InterceptorProvider<>(FetchRequest.class, new FetchRequestLoggerInterceptor()),
                new InterceptorProvider<>(FetchResponse.class,
                        new FetchResponseLoggerInterceptor(encryptor, deserializer, serializer)),
                new InterceptorProvider<>(ProduceRequest.class,
                        new ProduceLoggerInterceptor(encryptor, deserializer, serializer)),
                new InterceptorProvider<>(AbstractResponse.class, new ResponseLoggerInterceptor())
        );
    }
}
