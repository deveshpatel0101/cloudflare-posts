import './App.css';
import Posts from './Components/Posts/Posts';
import AddPost from './Components/AddPost/AddPost';
import React from 'react';
import { getPosts } from './controllers/posts';

class App extends React.Component {
  state = {
    posts: [],
  };

  async componentDidMount() {
    let res = await getPosts();
    this.setState({ posts: res });
  }

  updateState = (data) => {
    this.setState({ posts: data });
  };

  render() {
    return (
      <div className='App'>
        <AddPost updateState={this.updateState} />
        <Posts posts={this.state.posts} />
      </div>
    );
  }
}

export default App;
