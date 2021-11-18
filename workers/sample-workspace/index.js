addEventListener('fetch', event => {
  event.respondWith(handleRequest(event.request))
});

/**
 * Respond with hello worker text
 * @param {Request} request
 */
async function handleRequest(request) {
  if(request.method === 'OPTIONS') {
    return new Response('', {
      headers: {'Access-Control-Allow-Origin': '*', 'Access-Control-Allow-Headers': '*', 'Access-Control-Request-Headers': '*'}
    });
  }

  const path = request.url.split('workers.dev/')[1];
  if (path.toLowerCase().includes('posts')) {
    try {
      if(request.method === 'GET') {
        return await get(request);
      } else if(request.method === 'POST') {
        return await post(request);
      }
    } catch (error){
      return new Response(error, {
        headers: { 'content-type': 'application/json' },
        status: 500
      });
    }
  }

  return new Response('Route not available', {
    headers: { 'content-type': 'application/json' },
    status: 403
  });
}

/**
 * @param {Request} request
 */
const get = async (request) => {
  return new Response(await sample_kv.get('posts'), {
    headers: { 'content-type': 'application/json', 'Access-Control-Allow-Origin': '*' },
  });
}

/**
 * @param {Request} request
 */
const post = async (request) => {
  const posts = JSON.parse(await sample_kv.get('posts'));
  const post = await request.json();
  posts.push(post);
  await sample_kv.put('posts', JSON.stringify(posts));
  return new Response(JSON.stringify(posts), {
    headers: { 'content-type': 'application/json', 'Access-Control-Allow-Origin': '*' },
  });
}