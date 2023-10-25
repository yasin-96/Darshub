namespace pdfGen.common.Models.Course;
public class CourseWithChapters
{
    public string Id { get; set; } = string.Empty;
    public string Name { get; set; } = string.Empty;
    public string Description { get; set; } = string.Empty;
    public int Level { get; set; }
    public string Author { get; set; } = string.Empty;
    public DateTime Released { get; set; }
    public DateTime LastUpdate { get; set; }
    public string Duration { get; set; } = String.Empty;
    public bool IsContentMarkDown { get; set; }
    public List<ChapterWithSubChapters>? Content { get; set; } = new List<ChapterWithSubChapters>();
}

public class CourseWithChaptersInHTML
{
    public string Id { get; set; } = string.Empty;
    public string Name { get; set; } = string.Empty;
    public string Description { get; set; } = string.Empty;
    public int Level { get; set; }
    public string Author { get; set; } = string.Empty;
    public DateTime Released { get; set; }
    public DateTime LastUpdate { get; set; }
    public string Duration { get; set; } = String.Empty;
    public bool IsContentMarkDown { get; set; }
    public List<ChapterWithSubChaptersInHTML>? Content { get; set; } = new List<ChapterWithSubChaptersInHTML>();
}