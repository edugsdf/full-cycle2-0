//---------------------------------------------------
const express = require('express');
const mysql = require('mysql');
const app = express();

app.use(express.json()) // for parsing application/json
app.use(express.urlencoded({ extended: true })) // for parsing application/x-www-form-urlencoded

const port = 3000;

var strTextoNome = '';

const config = {
    connectionLimit : 10,
    host: 'srv_mysql',
    user: 'root',
    password: 'root',
    database: 'nodedb'
};

var pool  = mysql.createPool(config);
pool.getConnection((err) => {
    if(err){
      console.log('Error connecting to Db');
      return;
    }
    console.log('Connection established');
});

//----------POST
app.post('/add', (req, res) => {
    var numero = new Date().getTime();

    if (Object.keys(req.body).length !== 0) {
        pool.query('INSERT INTO people SET ?', {name: req.body.nome}, (error, results) => {
            if (error)
                throw error;
            console.log(`Gravado na base o nome ${req.body.nome}`);
        }).end();
    }

    res.send(`<h1> Cadastrado com sucesso!!! </h1>`+
            `<script>setTimeout(function(){ window.location.href = '/?nocache=${numero}' }, 1200);</script>`);

});


//----------HOME
app.get('/', (req, res) => {

    if (Object.keys(req.query).length !== 0) {
        pool.query(`SELECT id, name FROM people ORDER BY id DESC`,function (error, result) {
            if (error)
                throw error;
            strTextoNome = `<h3>Lista de nomes</h3>`;
            Object.keys(result).forEach(function (key) {
                var row = result[key];
                strTextoNome = strTextoNome +  `<p>NOME = ${row.name} </p> `;
                console.log(`Nome => ${row.name}`);
            });
        });
    }
    res.send(`<h1> Full Cycle Rocks! </h1>`+
            ` <hr />`+
            `<form action="/add" method="POST" id="form1">
                <input type="text" id="nome" name="nome" placeholder="Nome" required>
                &nbsp;<button type="submit" form="form1" value="Enviar">Enviar</button>
            </form>`+
            `<hr />`+
            `${strTextoNome}`+
            `<script> window.onload = function() {
                if(!window.location.hash) {
                    window.location = window.location + '#loaded';
                    window.location.reload();
                }
            }</script>`
    )
});


app.listen(port, () => {
    console.log(`Rodando na porta `+ port);
});