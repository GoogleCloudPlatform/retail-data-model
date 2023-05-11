def _gen_terraform_impl(ctx):
    # Generates a terraform script from a schema file.

    # This depends on the output of a BQ Schema build
    directoryToScan = ctx.files.srcs[0]

    input_files = []

    for t in ctx.attr.srcs:
        for f in t.files.to_list():
            if (ctx.attr.debug):
                print(f.path)
            input_files.append(f)

    output_files = []

    out_file = ctx.actions.declare_file("bigquery/main.tf")
    output_files.append(out_file)

    args = ctx.actions.args()

    # Local variable for the prefix
    args.add("--input", directoryToScan.path)
    args.add("--output", out_file.path)
    args.add("--project_id", ctx.attr.project_id)
    args.add("--location", ctx.attr.location)
    args.add("--access_owner", ctx.attr.access_owner)
    args.add("--access_owner_role", ctx.attr.access_owner_role)
    args.add("--access_reader_role", ctx.attr.access_reader_role)
    args.add("--access_reader_domain", ctx.attr.access_reader_domain)
    args.add("--expiration_time", ctx.attr.expiration_time)
    args.add("--dataset_id", ctx.attr.dataset_id)
    args.add("--table_id", ctx.attr.table_id)
    args.add("--time_partition_type", ctx.attr.time_partition_type)
    args.add("--time_partition_expiration_ms", ctx.attr.time_partition_expiration_ms)
    args.add("--time_partition_field", ctx.attr.time_partition_field)
    args.add("--debug", ctx.attr.debug)

    ctx.actions.run(
        inputs = [f],
        outputs = [out_file],
        arguments = [args],
        progress_message = "Generating terraform file %s from input directory: %s" % (out_file.path, directoryToScan.path),
        executable = ctx.executable.cmd,
    )

    return [DefaultInfo(files = depset(output_files))]

gen_terraform = rule(
    _gen_terraform_impl,
    attrs = {
        "debug": attr.bool(default = False),
        "project_id": attr.string(
            doc = "(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used.",
            default = "var.project",
        ),
        "location": attr.string(
            doc = "(Optional) The geographic location where the dataset should reside.",
            default = "var.location",
        ),
        "access_owner": attr.string(
            doc = "The email of the individual or team that owns the dataset",
            default = "var.owner",
            mandatory = True,
        ),
        "access_owner_role": attr.string(
            doc = "The name of the role to assign to the owner group.",
            default = "OWNER",
        ),
        "access_reader_role": attr.string(
            doc = "The name of the role to assigne to the reader group.",
            default = "READER",
        ),
        "access_reader_domain": attr.string(
            doc = "The domain name allowed reader access.",
            default = "var.bq_reader_domain",
        ),
        "expiration_time": attr.string(
            doc = "(Optional) The time when this table expires, in milliseconds since the epoch. If not present, the table will persist indefinitely. Expired tables will be deleted and their storage reclaimed.",
        ),
        "dataset_id": attr.string(
            doc = "(Required) A unique ID for this dataset, without the project name. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 1,024 characters.",
            mandatory = True,
        ),
        "table_id": attr.string(
            doc = "The name of table, this is an overide only function and should only be used on single tables.",
        ),
        "time_partition_type": attr.string(
            doc = "(Required) The supported types are DAY, HOUR, MONTH, and YEAR, which will generate one partition per day, hour, month, and year, respectively.",
            default = "DAY",
        ),
        "time_partition_expiration_ms": attr.string(
            doc = "Optional) Number of milliseconds for which to keep the storage for a partition.",
        ),
        "time_partition_field": attr.string(
            doc = "(Optional) The field used to determine how to create a time-based partition. If time-based partitioning is enabled without this value, the table is partitioned based on the load time.",
        ),
        "cmd": attr.label(
            allow_files = True,
            executable = True,
            cfg = "exec",
        ),
        "srcs": attr.label_list(
            allow_files = True,
            doc = "Source files to write the terraform scritps from.",
        ),
    },
    doc = "Writes terraform scripts to the schema directory",
    executable = False,
)
