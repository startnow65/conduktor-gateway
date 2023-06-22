FROM maven:3.8.5-openjdk-17 as builder
WORKDIR /app

COPY pom.xml ./
COPY gateway-core/pom.xml ./gateway-core/pom.xml
COPY interceptor-framework/pom.xml ./interceptor-framework/pom.xml
COPY logger-interceptor/pom.xml ./logger-interceptor/pom.xml
COPY gateway-encryption/pom.xml ./gateway-encryption/pom.xml
RUN mvn --fail-never dependency:go-offline

COPY . .
RUN mvn clean package

FROM ubuntu/jre:17-22.04_2
WORKDIR /app
COPY --from=builder /app/gateway-core/target/gateway-core-0.5.0-SNAPSHOT.jar ./
COPY --from=builder /app/logger-interceptor/target/logger-interceptor-0.5.0-SNAPSHOT.jar ./plugins/
COPY --from=builder /app/gateway-encryption/target/gateway-encryption-0.5.0-SNAPSHOT.jar ./plugins/
COPY --from=builder /app/gateway-core/config ./
ENTRYPOINT [ "/opt/java/bin/java", "--class-path", "/app/*:/app/plugins/*", "io.conduktor.gateway.Bootstrap" ]
