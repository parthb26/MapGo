import React, { useEffect, useState } from "react";
import axios from "axios";

// Traffic Data Interface
interface Traffic {
  location: string;
  status: string;
  estimated_time: string;
}

const TrafficComponent: React.FC = () => {
  const [trafficData, setTrafficData] = useState<Traffic | null>(null);

  // Fetch traffic data from the API
  useEffect(() => {
    const fetchTrafficData = async () => {
      try {
        const response = await axios.get("http://localhost:8080/traffic");
        setTrafficData(response.data); // Set the response data to state
      } catch (error) {
        console.error("Error fetching traffic data:", error);
      }
    };

    fetchTrafficData();
  }, []); // Empty dependency array ensures this runs once when the component mounts

  return (
    <div>
      <h1>Traffic Information</h1>
      {trafficData ? (
        <div>
          <p><strong>Location:</strong> {trafficData.location}</p>
          <p><strong>Status:</strong> {trafficData.status}</p>
          <p><strong>Estimated Time:</strong> {trafficData.estimated_time}</p>
        </div>
      ) : (
        <p>Loading traffic data...</p>
      )}
    </div>
  );
};

export default TrafficComponent;
