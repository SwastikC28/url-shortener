import { Route, Routes } from 'react-router';
import Home from './pages/Home';
import Analytics from './pages/Analytics';

function App() {
  return (
    <div className="d-flex flex-column site-container">
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path='/analytics/:shortURL' element={<Analytics/>}/>
      </Routes>
    </div>
  );
}

export default App;
