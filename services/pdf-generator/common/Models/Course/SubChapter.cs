using HtmlAgilityPack;
using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;

namespace pdfGen.common.Models.Course;
[BsonIgnoreExtraElements]
public class SubChapter
{
    [BsonId]
    [BsonRepresentation(BsonType.ObjectId)]
    public string Id { get; set; } = string.Empty;

    [BsonElement("title")] public string Title { get; set; }
    [BsonElement("content")] public string Content { get; set; }
    [BsonElement("listing")] public List<string> Listing { get; set; }
}


public class SubChapterInHtml
{
    public string Id { get; set; } = string.Empty;
    public HtmlDocument Content { get; set; }
    public List<HtmlDocument> Listing { get; set; }
}