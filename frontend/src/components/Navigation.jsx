import { Link } from 'react-router-dom'

function Navigation() {
  return (
    <nav style={{ padding: '1rem', borderBottom: '1px solid #ccc', marginBottom: '2rem' }}>
        <Link to="/jobs" style={{ marginRight: '1rem' }}>Jobs</Link>
        <Link to="/create">Create Job</Link>
    </nav>
  )
}

export default Navigation
