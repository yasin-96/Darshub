using MongoDB.Bson.Serialization.Attributes;

namespace pdfGen.common.Models.Course;
[BsonIgnoreExtraElements]
public class Category 
{
    [BsonElement("name")]
    public string Name { get; set; }

    [BsonElement("description")]
    public string Description { get; set; }

    [BsonElement("skills")]
    public string Skills { get; set; }
}