GRPC_VERSION = "1.61.0"
PROTOBUF_VERSION = "3.25.3"

JAVA_DEPS = [
    "info.picocli:picocli:4.7.5",
    "info.picocli:picocli-codegen:4.7.5",
    "org.apache.commons:commons-lang3:3.13.0",
    "com.google.protobuf:protobuf-java-util:%s" % PROTOBUF_VERSION,
    "io.grpc:grpc-core:%s" % GRPC_VERSION,
    "io.grpc:grpc-googleapis:%s" % GRPC_VERSION,
    "io.grpc:grpc-netty-shaded:%s" % GRPC_VERSION,
    "io.grpc:grpc-protobuf:%s" % GRPC_VERSION,
    "io.grpc:grpc-stub:%s" % GRPC_VERSION,
    "io.grpc:grpc-testing:%s" % GRPC_VERSION,
    "org.apache.tomcat:annotations-api:6.0.53",
    "org.apache.logging.log4j:log4j-api:2.21.0",
    "org.apache.logging.log4j:log4j-core:2.21.0",
    "com.google.protobuf:protoc:%s" % PROTOBUF_VERSION,
    "io.netty:netty-all:4.1.100.Final",
]

JUNIT_JUPITER_VERSION = "5.10.1"
JUNIT_PLATFORM_VERSION = "1.10.2"

TEST_DEPS = [
    "org.junit.platform:junit-platform-launcher:%s" % JUNIT_PLATFORM_VERSION,
    "org.junit.platform:junit-platform-reporting:%s" % JUNIT_PLATFORM_VERSION,
    "org.junit.platform:junit-platform-console:%s" % JUNIT_PLATFORM_VERSION,
    "org.junit.jupiter:junit-jupiter-api:%s" % JUNIT_JUPITER_VERSION,
    "org.junit.jupiter:junit-jupiter-params:%s" % JUNIT_JUPITER_VERSION,
    "org.junit.jupiter:junit-jupiter-engine:%s" % JUNIT_JUPITER_VERSION,
    "io.javalin:javalin-testtools:5.6.3",
    "org.jboss.weld:weld-junit5:4.0.1.Final",
    "org.jboss.weld:weld-junit-common:4.0.1.Final",
]
