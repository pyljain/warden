const express = require('express')
const app = express()
const port = process.env["PORT"]

app.get("/", (req, res) => {
    res.end(process.env['OUTPUT'])
})

app.listen(port, () => {
    console.log("Listening on port", port)
})