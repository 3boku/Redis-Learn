import { createClient } from "redis"
import express from "express"

const client = createClient ({
  url : "redis url"
});

client.on("error", function(err) {
    throw err;
});
client.connect().then();

const app = express()
app.use(express.json())
app.use(express.urlencoded({extended: true}))
const port = 3000

app.get('/', async(req, res) => {
    const cache = await client.get("posts");
    if(cache) return res.json(JSON.parse(cache))
    const result = await fetch('https://jsonplaceholder.typicode.com/posts/1')
    const json = await result.json();
    await client.set("posts", JSON.stringify(json));
    res.send(json)
})

app.get("/posts/:id", async(req, res) => {
    const id = req.params.id;
    if(!id) res.sendStatus(400);
    const cache = await client.get(`v1@posts/${id}`);
    if(cache) return res.json(JSON.parse(cache))
    const result = await fetch('https://jsonplaceholder.typicode.com/posts/1')
    const json = await result.json();
    await client.set(`v1@posts/${id}`, JSON.stringify(json));

    res.send(json)
})

app.listen(port, () => {
    console.log(`Example app listening on port ${port}`)
})
