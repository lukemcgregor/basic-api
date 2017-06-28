var express    = require('express');
var app        = express();
var bodyParser = require('body-parser');
var xml = require('xml');

require('body-parser-xml')(bodyParser);

app.use(bodyParser.json());
app.use(bodyParser.xml());

var port = process.env.PORT || 8080;

app.get('/add/:x/to/:y', (req, res) => {
    res.json({ sum: parseInt(req.params.x) + parseInt(req.params.y) });
});

app.post('/add', (req, res) => {
    var sum = req.body.numbers.value.map(parseInt).reduce((a, b) => a + b, 0);
    res.format({
        xml: () => {
            res.send(xml({sum:sum}));
        },
        json: () => {
            res.json({sum:sum});
        }
    })
});

app.listen(port);
