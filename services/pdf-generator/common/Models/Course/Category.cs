using pdfGen.common.Models.Base;

namespace pdfGen.common.Models.Course;
public class Category : BaseContent
{
    [BsonElement("name")]
    public string Name { get; set; }

    [BsonElement("description")]
    public string Description { get; set; }

    [BsonElement("skills")]
    public string Skills { get; set; }
}