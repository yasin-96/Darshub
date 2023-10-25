using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;

namespace pdfGen.common.Models.Course;

[BsonIgnoreExtraElements]
public class Chapter
{
    [BsonId]
    [BsonRepresentation(BsonType.ObjectId)]
    public string Id { get; set; } = string.Empty;

    [BsonElement("title")] public string Title { get; set; }

    [BsonElement("description")] public string Description { get; set; }

    [BsonElement("skills")] public string Skills { get; set; }

    [BsonElement("subchapters")] public List<ObjectId> SubChapterReferences { get; set; }
}

public class ChapterWithSubChapters
{
    public string Id { get; set; } = string.Empty;

    public string Title { get; set; } = string.Empty;

    public string Description { get; set; } = string.Empty;

    public string Skills { get; set; } = string.Empty;

    public List<SubChapter> SubChapterReferences { get; set; } = new List<SubChapter>();
}


public class ChapterWithSubChaptersInHTML
{
    public string Id { get; set; } = string.Empty;

    public string Name { get; set; } = string.Empty;

    public string Description { get; set; } = string.Empty;

    public string Skills { get; set; } = string.Empty;

    public List<SubChapterInHtml> SubChapterReferences { get; set; } = new List<SubChapterInHtml>();
}