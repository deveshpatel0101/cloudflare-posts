export const getPosts = async () => {
  let res = await fetch(
    'https://sample-workspace.dpatel13599.workers.dev/posts/'
  );
  return await res.json();
};

export const putPost = async (data) => {
  try {
    let res = await fetch(
      'https://sample-workspace.dpatel13599.workers.dev/posts/',
      {
        method: 'POST',
        body: JSON.stringify(data),
        headers: { 'Content-Type': 'application/json' },
      }
    );
    return await res.json();
  } catch (error) {
  }
};
