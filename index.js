
const express = require('express');
const os = require('os');

const app = express();
const port = 3000

// Disable client caching to make different pods more visible in testing.
app.use((req, res, next) => {
  res.set('Cache-Control', 'no-store');
  res.header('Expires', '-1');
  res.header('Pragma', 'no-cache');  next();
});

app.get('/', (req, res) => {
  res.send(`Host: ${os.hostname}`);
});

app.listen(port, () => {
  console.log(`App listening at on port: ${port}`)
});
