import React from "react";
import Navigation from "../components/Navigation";

const Workouts: React.FC = () => {
  return (
    <div className="min-h-screen bg-gradient-to-br from-green-50 to-emerald-100">
      <Navigation />
      <div className="container mx-auto px-4 py-16">
        <div className="text-center">
          <h1 className="text-4xl font-bold text-gray-900 mb-6">
            Workout Plans
          </h1>
          <p className="text-xl text-gray-600 mb-8">
            Browse and create your personalized workout routines
          </p>
        </div>

        <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div className="bg-white p-6 rounded-xl shadow-lg">
            <h3 className="text-xl font-semibold text-gray-900 mb-2">
              Beginner Plan
            </h3>
            <p className="text-gray-600 mb-4">
              Perfect for those just starting their fitness journey
            </p>
            <button className="bg-green-600 hover:bg-green-700 text-white font-semibold py-2 px-4 rounded-lg transition-colors duration-200">
              View Plan
            </button>
          </div>

          <div className="bg-white p-6 rounded-xl shadow-lg">
            <h3 className="text-xl font-semibold text-gray-900 mb-2">
              Intermediate Plan
            </h3>
            <p className="text-gray-600 mb-4">
              For those with some fitness experience
            </p>
            <button className="bg-green-600 hover:bg-green-700 text-white font-semibold py-2 px-4 rounded-lg transition-colors duration-200">
              View Plan
            </button>
          </div>

          <div className="bg-white p-6 rounded-xl shadow-lg">
            <h3 className="text-xl font-semibold text-gray-900 mb-2">
              Advanced Plan
            </h3>
            <p className="text-gray-600 mb-4">
              Challenging workouts for fitness enthusiasts
            </p>
            <button className="bg-green-600 hover:bg-green-700 text-white font-semibold py-2 px-4 rounded-lg transition-colors duration-200">
              View Plan
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Workouts;
