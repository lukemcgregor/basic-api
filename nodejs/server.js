var express    = require('express');
var app        = express();
var bodyParser = require('body-parser');
var xml = require('xml');

require('body-parser-xml')(bodyParser);

app.use(bodyParser.json());
app.use(bodyParser.xml());

var port = process.env.PORT || 8080;
var router = express.Router();

router.get('/add/:x/to/:y', function(req, res) {
    res.json({ sum: parseInt(req.params.x) + parseInt(req.params.y) });
});

router.post('/add', function(req, res) {
    var sum = req.body.numbers.value.reduce((a, b) => parseInt(a) + parseInt(b), 0);
    res.format({
        xml: () => {
            res.send(xml({sum:sum}));
        },
        json: () => {
            res.json({sum:sum});
        }
    })
});

app.use('/', router);
app.listen(port);
