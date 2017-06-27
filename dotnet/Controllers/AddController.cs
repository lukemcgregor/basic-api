using System;
using System.Linq;
using BasicApi.Models;
using Microsoft.AspNetCore.Mvc;

namespace BasicApi.Controllers
{
    public class AddController : Controller
    {
        [HttpGet]
        [Route("add/{x}/to/{y}")]
        public IActionResult Get(int x, int y)
        {
            return Ok(new { sum = x + y});
        }

        [HttpPost]
        [Route("add")]
        public IActionResult Post([FromBody]Numbers numbers)
        {
            return Ok(new { sum = numbers.Values.Sum() });
        }
    }
}
