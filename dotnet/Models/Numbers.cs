using System.Xml.Serialization;

namespace BasicApi.Models
{
    [XmlRoot(ElementName = "numbers")]
    public class Numbers
    {   
        [XmlElement("value")]     
        public int[] Values { get; set; }
    }
}