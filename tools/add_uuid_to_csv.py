import getopt
import sys
import uuid


def print_help():
    print ('transform_for_bq -i <input_file> -o <output_file>')


def main(argv):
    input_file = ''
    output_file = ''
    opts, args = getopt.getopt(argv, "hi:o:", ["inFile=", "outFile="])
    for opt, arg in opts:
        if opt == '-h':
            print_help()
            sys.exit(0)
        elif opt in ("-i", "--inFile"):
            input_file = arg
        elif opt in ("-o", "--outFile"):
            output_file = arg

    if len(input_file) == 0 or len(output_file) == 0:
        print_help()
        print("Input and output files are required")
        sys.exit(1)

    print("Reading from {0} and writing to: {1}".format(input_file, output_file))

    in_ = open(input_file, 'r')
    out_ = open(output_file, 'w')
    in_l = in_.readlines()
    count = 0
    for line in in_l:
        u_ = uuid.uuid4()
        if count == 0:
            out = "id," + line
        else:
            out = str(u_) + "," + line
        count = count + 1
        out_.write(out)


if __name__ == "__main__":
    main(sys.argv[1:])