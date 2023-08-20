
namespace pdfGen.common.Models.Course;
public class Course
{
    [BsonElement("_id")]
    public string Id { get; set; }

    [BsonElement("name")]
    public string Name { get; set; }

    [BsonElement("description")]
    public string Description { get; set; }

    [BsonElement("level")]
    public int Level { get; set; }

    [BsonElement("chapters")]
    public List<string> ChapterReferences { get; set; }

    [BsonElement("author")]
    public string Author { get; set; }

    [BsonElement("released")]
    public long Released { get; set; }

    [BsonElement("lastUpdate")]
    public long LastUpdate { get; set; }
}