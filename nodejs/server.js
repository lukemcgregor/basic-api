var app = require('express')()
var xml = require('xml')

var bodyParser = require('body-parser')
require('body-parser-xml')(bodyParser)
app.use(bodyParser.json())
app.use(bodyParser.xml())

app.get('/add/:x/to/:y', (req, res) => {
    res.json({ sum: parseInt(req.params.x) + parseInt(req.params.y) })
})

app.post('/add', (req, res) => {
    var sum = req.body.numbers.value.map(parseInt).reduce((a, b) => a + b)
    res.format({
        xml: () => res.send(xml({sum:sum})),
        json: () => res.json({sum:sum})
    })
})

app.listen(8080)