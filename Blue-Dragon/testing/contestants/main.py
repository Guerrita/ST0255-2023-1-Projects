from fastapi import FastAPI
from fastapi.staticfiles import StaticFiles
from fastapi.responses import HTMLResponse
from fastapi import Request

app = FastAPI()

# Serve static files from the "public" directory
app.mount("/public", StaticFiles(directory="public"), name="public")

# Serve the "index.html" file as the root URL
@app.get("/", response_class=HTMLResponse)
async def read_root(request: Request):
    return await request.app.mount("/public").get_path("/index.html")

# Start the server
if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="localhost", port=3000)
