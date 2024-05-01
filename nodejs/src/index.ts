import express, { Request, Response } from "express";
import { NodeTracerProvider } from "@opentelemetry/node";
import {
  ConsoleSpanExporter,
  SimpleSpanProcessor,
} from "@opentelemetry/tracing";
import { ZipkinExporter } from "@opentelemetry/exporter-zipkin";
import path from "node:path";

const app = express();
app.use(express.json())
const port = 3000;

const provider = new NodeTracerProvider();
const consoleExporter = new ConsoleSpanExporter();
const spanProcessor = new SimpleSpanProcessor(consoleExporter);
provider.addSpanProcessor(spanProcessor);
provider.register();

const zipkinExporter = new ZipkinExporter({
  url: "http://localhost:9411/api/v2/spans",
  serviceName: "course-service",
});

const zipkinProcessor = new SimpleSpanProcessor(zipkinExporter);
provider.addSpanProcessor(zipkinProcessor);

app.get("/", async function (req: Request, res: Response): Promise<void> {
  res.type("json");

  // Delay 50s
  await new Promise((resolve) => setTimeout(resolve, 50));
  //

  // Connect to SQLITE Database
    const sqlite3 = require("sqlite3").verbose();

    let db = new sqlite3.Database(
      path.resolve(__dirname, "db.sqlite3"),
      // sqlite3.OPEN_READWRITE,
      (err: any) => {
        if (err) {
          console.log("Error");
          console.log(err);
          console.error(err.message);
        }

        console.log("Connected to the  database.");
      }
    );
  //

  // SELECT from Database
  db.all(`SELECT * FROM courses`, [], (err: any, rows: any) => {
    console.log({ err, rows });
    
    res.status(200).json({ rows });
  });
  //
});

app.listen(port, () => {
  console.log(`Listening at http://localhost:${port}`);
});
