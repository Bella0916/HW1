
import './styles/App.css';
import { MOONNavbar } from './components/MOONNavbar';
import { ProductContent } from './components/ProductContent';
import { CustomerContent } from './components/CustomerContent';
import { OrderContent } from './components/OrderContent';
import { ItemContent } from './components/ItemContent';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

const NotFoundContent = () => {
  return (
      <div className="App">
          NotFoundContent
      </div>
  );
};

  function App() {
    return (
      <> 
        <MOONNavbar />
        <Router>
          <Routes>
              <Route exact path="/products" element={<ProductContent />} />
              <Route exact path="/customers" element={<CustomerContent />} />
              <Route exact path="/orders" element={<OrderContent />} />
              <Route exact path="/items" element={<ItemContent/>} />
              <Route path="*" element={<NotFoundContent />} />
          </Routes>
        </Router>
      </>
      
    );
  }

export default App;
