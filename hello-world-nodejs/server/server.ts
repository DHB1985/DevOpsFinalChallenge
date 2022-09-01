import express from "express";

const PORT = process.env.PORT || 3000;
// Create a new express app instance
const app: express.Application = express();
app.get("/hello", function (req, res) {
  res.send("Hello World!");
});
// app.listen(3000, "127.0.0.1", function () { //Se elimina el segundo argumento
app.listen(PORT, function () {
  console.log("App is listening on port ", PORT, "!");
});
