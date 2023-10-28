using pdfGen.common.Models.Course;
using Microsoft.Extensions.Logging;
using Microsoft.Extensions.Options;
using MongoDB.Driver;
using service.mongo.settings;

namespace service.mongo;
public class MongoService
{
    private readonly ILogger<MongoService> _logger;
    private readonly MongoClient _mongoClient;
    private readonly IMongoDatabase _mongoDB;

    public MongoService(ILogger<MongoService> logger, IOptions<MongoSettings> mongoSettings)
    {
        _logger = logger;
        _mongoClient = new MongoClient(mongoSettings.Value.URL);
        _mongoDB = _mongoClient.GetDatabase(mongoSettings.Value.Database);
    }

    public async Task<bool> IsCourseAvailable(string courseId)
    {
        try
        {
            var filter = Builders<Course>.Filter.Eq<string>(doc => doc.Id, courseId);
            var mdbResult = await _mongoDB.GetCollection<Course>("course").Find(Builders<Course>.Filter.Empty)
                .SingleOrDefaultAsync();

            if (mdbResult is null)
            {
                return false;
            }

            return true;
        }
        catch (Exception e)
        {
            Console.WriteLine(e);
            return false;
        }
    }

    public async Task<CourseWithChapters> GetCourseById(string courseId)
    {
        if (string.IsNullOrEmpty(courseId))
        {
            _logger.LogError("MDB Error: Missing information for requesting data.");
            return default;
        }
        
        
        CourseWithChapters? mdResponse = null;
        
        try
        {
            var matchCourse = Builders<Course>.Filter.Eq(doc => doc.Id, courseId);
            var courseResult = await _mongoDB.GetCollection<Course>("course")
                .Find<Course>(matchCourse).SingleOrDefaultAsync();
            
            if (courseResult is null)
            {
                _logger.LogError("MDB Error: Course not found");
                return default;
            }

            mdResponse = new CourseWithChapters()
            {
                Id = courseResult.Id,
                Author = courseResult.Author,
                Description = courseResult.Description,
                Duration = courseResult.Duration,
                IsContentMarkDown = courseResult.IsContentMarkDown,
                Level = courseResult.Level,
                LastUpdate = courseResult.LastUpdate,
                Released = courseResult.Released,
            };
                
            var matchChapter = Builders<Chapter>.Filter.In("_id", courseResult.ChapterReferences);
            var chapterResults = await _mongoDB.GetCollection<Chapter>("chapter")
                .Find<Chapter>(matchChapter).ToListAsync();

            if (chapterResults is null || chapterResults.Count <=0)
            {
                _logger.LogError("MDB Error: Content for course not found");
                return default;
            }

            var chapterList = new List<Chapter>(); 
            foreach (var chapter in chapterResults)
            {
                var chapWithSubChapter = new ChapterWithSubChapters()
                {
                    Id = chapter.Id,
                    Description = chapter.Description,
                    Title = chapter.Title,
                    Skills = chapter.Skills,
                };
                
                if (chapter.SubChapterReferences.Count >= 0 || chapter.SubChapterReferences is not null)
                {
                    var matchSubChapters = Builders<SubChapter>.Filter.In("_id", chapter.SubChapterReferences);
                    var subChapterResult = await _mongoDB.GetCollection<SubChapter>("subchapter")
                        .Find<SubChapter>(matchSubChapters).ToListAsync();
                    if (subChapterResult is not null && subChapterResult.Count > 0)
                    {
                        chapWithSubChapter.SubChapterReferences.AddRange(subChapterResult);
                    } 
                }
                mdResponse.Content.Add(chapWithSubChapter);
            }

            return mdResponse;
        }
        catch (Exception e)
        {
            _logger.LogError(e, "GetCourseById General Ex: {emsg}", e.Message);
            return default;
        }
    }
}