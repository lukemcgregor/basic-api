using Microsoft.AspNetCore.Hosting;

namespace BasicApi
{
    public class Program
    {
        public static void Main(string[] args)
        {
            var host = new WebHostBuilder()
                .UseKestrel()
                .UseStartup<Startup>()
                .UseUrls("http://localhost:8080")
                .Build();

            host.Run();
        }
    }
}
