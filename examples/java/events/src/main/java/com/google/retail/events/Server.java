package com.google.retail.events;

import com.google.retail.events.service.EventRecordsService;
import com.google.retail.events.service.EventsService;
import com.google.retail.events.service.StatusService;
import io.grpc.ServerBuilder;
import io.grpc.ServerServiceDefinition;
import picocli.CommandLine;
import picocli.CommandLine.Command;

import java.io.File;
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.Callable;
import java.util.logging.Logger;

@Command(name = "server", mixinStandardHelpOptions = true, version = "Events Server 1.0",
    description = "Starts and stops the the GRPC server for events")
public class Server implements Callable<Integer> {
  private static final Logger LOG = Logger.getLogger(Logger.GLOBAL_LOGGER_NAME);

  @CommandLine.Option(names = {"-p", "--port"}, description = "Port number, default: 8080", defaultValue = "8080")
  private int port;

  @CommandLine.Option(names = {"-f", "--featureFile"}, description = "A url to a feature file")
  private File featureFile;

  @Override
  public Integer call() throws Exception {
    LOG.info(String.format("Starting server at port %d", port));

    var eventDescriptionService = new EventRecordsService().bindService();
    var eventRecordService = new EventRecordsService().bindService();
    var eventsService = new EventsService().bindService();
    var statusService = new StatusService().bindService();

    List<ServerServiceDefinition> availableServices = new ArrayList<>();
    availableServices.add(eventDescriptionService);
    availableServices.add(eventRecordService);
    availableServices.add(eventsService);
    availableServices.add(statusService);

    var service = ServerBuilder
        .forPort(port).addServices(availableServices).build();

    var server = service.start();

    Runtime.getRuntime().addShutdownHook(new Thread() {
      public void run() {
        System.err.println("*** Shutting down gRPC server since JVM is shutting down");
        server.shutdown();
        System.out.println("gRPC Server Shutdown.");
      }
    });

    try {
      server.awaitTermination();
    } catch (InterruptedException e) {
      LOG.warning("Thread interrupted: " + e.getMessage());
      Thread.currentThread().interrupt();
    }
    return 0;
  }

  public static void main(String[] args) {
    int exitCode = new CommandLine(new Server()).execute(args);
    System.exit(exitCode);
  }

}
