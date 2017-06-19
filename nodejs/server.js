var express    = require('express');
var app        = express();
//var bodyParser = require('body-parser');

//app.use(bodyParser.urlencoded({ extended: true }));
//app.use(bodyParser.json());

var port = process.env.PORT || 8080;
var router = express.Router();

router.get('/add/:x/to/:y', function(req, res) {
    res.json({ sum: parseInt(req.params.x) + parseInt(req.params.y) });
});

app.use('/', router);
app.listen(port);
