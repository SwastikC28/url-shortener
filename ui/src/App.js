import { Route, Routes } from 'react-router';
import Home from './pages/Home';

function App() {
  return (
    <div className="d-flex flex-column site-container">
      <Routes>
        <Route path="/" element={<Home />} />
      </Routes>
    </div>
  );
}

export default App;
