# Proto definitions used throughout the project
# Go dependencies are now handled by MODULE.bazel

def go_dependencies():
    """ Go dependencies are now handled by MODULE.bazel - this function is kept for compatibility """
    pass

BQ_PROTOS = [
    "@protoc_gen_bq_schema//api:bq_schema_proto_lib",
]

COMMON_API_PROTOS = [
    "@com_google_protobuf//:timestamp_proto",
    "@com_google_protobuf//:struct_proto",
    "@com_google_protobuf//:any_proto",
    "@com_google_protobuf//:api_proto",
    "@com_google_protobuf//:empty_proto",
    "@com_google_protobuf//:field_mask_proto",
    "@com_google_protobuf//:descriptor_proto",
    "@com_google_googleapis//google/api:annotations_proto",
]

WELL_KNOWN_TYPES = [
    "@org_golang_google_protobuf//types/known/anypb:go_default_library",
    "@org_golang_google_protobuf//types/known/apipb:go_default_library",
    "@org_golang_google_protobuf//types/known/durationpb:go_default_library",
    "@org_golang_google_protobuf//types/known/emptypb:go_default_library",
    "@org_golang_google_protobuf//types/known/fieldmaskpb:go_default_library",
    "@org_golang_google_protobuf//types/known/sourcecontextpb:go_default_library",
    "@org_golang_google_protobuf//types/known/structpb:go_default_library",
    "@org_golang_google_protobuf//types/known/timestamppb:go_default_library",
    "@org_golang_google_protobuf//types/known/typepb:go_default_library",
    "@org_golang_google_protobuf//types/known/wrapperspb:go_default_library",
]
