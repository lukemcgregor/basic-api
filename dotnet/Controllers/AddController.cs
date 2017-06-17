using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
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
        public IActionResult Post([FromBody]int[] numbers)
        {
            return Ok(new { sum = numbers.Sum() });
        }
    }
}
