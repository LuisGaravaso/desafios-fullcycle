const express = require("express");
const mysql = require("mysql2/promise");
const { v4: uuidv4 } = require("uuid");

// Random name generator
function generateRandomName() {
  const firstNames = [
    "Gabriela", "Lucas", "Maria", "João", "Carla", 
    "Pedro", "Beatriz", "Mateus", "Ana", "Ricardo"
  ];
  const lastNames = [
    "Silva", "Santos", "Oliveira", "Souza", "Pereira", 
    "Costa", "Almeida", "Ferreira", "Rodrigues", "Gomes"
  ];

  const firstName = firstNames[Math.floor(Math.random() * firstNames.length)];
  const lastName = lastNames[Math.floor(Math.random() * lastNames.length)];
  return `${firstName} ${lastName}`;
}

// Database connection setup
const dbConfig = {
  host: "db",
  user: "root",
  password: "root",
  database: "fullcycle",
};

async function createConnection() {
  return await mysql.createPool(dbConfig);
}

const app = express();

// Route handlers
app.get("/people", async (req, res) => {
  const db = await createConnection();

  try {
    // Insert a random person into the database
    const randomPerson = {
      id: uuidv4(),
      name: generateRandomName(),
    };

    await db.execute("INSERT INTO people (id, name) VALUES (?, ?)", [
      randomPerson.id,
      randomPerson.name,
    ]);

    // Fetch all people from the database
    const [rows] = await db.execute("SELECT id, name FROM people");

    // Construct the HTML response
    let responseHTML = `
      <h1>Full Cycle Rocks!</h1>
      <p>Pessoas que visitaram essa página:</p>
    `;
    rows.forEach((person) => {
      responseHTML += `<p>${person.name}</p>`;
    });

    res.setHeader("Content-Type", "text/html");
    res.status(200).send(responseHTML);
  } catch (err) {
    console.error(err);
    res.status(500).send("Internal Server Error");
  }
});

// Start the server
const PORT = 9000;
app.listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});
