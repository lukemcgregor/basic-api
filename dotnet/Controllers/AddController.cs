using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using System.Xml.Serialization;
using Microsoft.AspNetCore.Mvc;

namespace basic_api_dotnet.Controllers
{
    public class AddController : Controller
    {
        [HttpGet]
        [Route("add/{x}/to/{y}")]
        public IActionResult Get(int x, int y)
        {
            return Ok(x + y);
        }

        [HttpPost]
        [Route("add")]
        public IActionResult Post([FromBody]Numbers numbers)
        {
            return Ok(new { sum = numbers.Values.Sum() });
        }
    }

    [XmlRoot(ElementName = "numbers")]
    public class Numbers
    {   
        [XmlElement("value")]     
        public int[] Values { get; set; }
    }
}
