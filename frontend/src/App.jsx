import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom'
import Navigation from './components/Navigation'
import JobsList from './pages/JobsList'
import CreateJob from './pages/CreateJob'

function App() {
  return (
    <BrowserRouter>
      <Navigation />
      <div style={{ padding: '1rem' }}>
        <Routes>
          <Route path="/" element={<Navigate to="/jobs" replace />} />
          <Route path="/jobs" element={<JobsList />} />
          <Route path="/create" element={<CreateJob />} />
        </Routes>
      </div>
    </BrowserRouter>
  )
}

export default App
