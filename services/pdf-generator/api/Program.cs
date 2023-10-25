using pdfGen.service.parser;
using pdfGen.service.pdf;
using service.mongo;
using service.mongo.settings;

var builder = WebApplication.CreateBuilder(args);

builder.Configuration.Sources.Clear();

IHostEnvironment env = builder.Environment;

builder.Configuration
    .AddJsonFile("appsettings.json", optional: false, reloadOnChange: true)
    .AddJsonFile($"appsettings.{env.EnvironmentName}.json", true, true);

builder.Services.Configure<MongoSettings>(
    builder.Configuration.GetSection(
        key: nameof(MongoSettings)
    )
);

// Add services to the container.
builder.Services
    .AddSingleton<MongoService>()
    .AddSingleton<PdfService>();


builder.Services.AddControllers();
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();

app.UseAuthorization();

app.MapControllers();

app.Run();