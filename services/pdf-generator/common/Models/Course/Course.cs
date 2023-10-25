using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;

namespace pdfGen.common.Models.Course;

[BsonIgnoreExtraElements]
public class Course
{
    [BsonId]
    [BsonRepresentation(BsonType.ObjectId)]
    public string Id { get; set; } = string.Empty;

    [BsonElement("name")] public string Name { get; set; } = string.Empty;

    [BsonElement("description")] public string Description { get; set; } = string.Empty;

    [BsonElement("level")] public int Level { get; set; }

    [BsonElement("chapters")] public List<ObjectId> ChapterReferences { get; set; } = new List<ObjectId>();

    [BsonElement("author")] public string Author { get; set; } = string.Empty;

    [BsonElement("released")] public DateTime Released { get; set; }

    [BsonElement("lastUpdate")] public DateTime LastUpdate { get; set; }

    [BsonElement("duration")] public string Duration { get; set; } =string.Empty;

    [BsonElement("markdown")] public bool IsContentMarkDown { get; set; }
}

public class CourseInHTML
{
    public CourseHeader Header { get; set; }
    public CourseMisc Misc { get; set; }
    public List<ChapterWithSubChapters>? Content { get; set; } = new List<ChapterWithSubChapters>();
}

public class CourseHeader {
    public string Id { get; set; } = string.Empty;
    public string Name { get; set; } = string.Empty;
    public string Description { get; set; } = string.Empty;
    public string Author { get; set; } = string.Empty;
}

public class CourseMisc
{
    public int Level { get; set; } 
    public DateTime Released { get; set; }
    public DateTime LastUpdate { get; set; }
    public string Duration { get; set; } = String.Empty;
}