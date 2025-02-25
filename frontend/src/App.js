import { useState } from "react";

export default function TaxCalculatorForm() {
  const [formData, setFormData] = useState({
    grossIncome: "",
    otherIncome: "",
    employerPf: "",
    isEmployerPfSeparate: false,
    employeePf: "",
    employerNps: false,
    basicIncome: "",
    contributionPercentage: "",
    oldRegime: false,
    employeeNps: "",
    savingsInterest: "",
    capitalGain: false,
    shortTermGain: "",
    longTermGain: "",
    dividendEarned: ""
  });
  
  const [taxResult, setTaxResult] = useState(null);
  const [showSavingsMessage, setShowSavingsMessage] = useState(false);

  const handleChange = (e) => {
    const { name, value, type, checked } = e.target;
    setFormData({
      ...formData,
      [name]: type === "radio" ? value === "true" : value,
    });
  };

  const handleSavingsBlur = () => {
    if (formData.savingsInterest !== "") {
      setShowSavingsMessage(true);
    } else {
      setShowSavingsMessage(false);
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const response = await fetch("http://localhost:8080/calculate-tax", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(formData),
    });
    const data = await response.json();
    setTaxResult(data);
  };

  return (
    <div className="max-w-lg mx-auto p-6 bg-blue-500 rounded-lg shadow-lg">
      <h2 className="text-2xl font-bold text-white mb-4 text-center">Income Tax Calculator</h2>
      <form onSubmit={handleSubmit} className="space-y-4">
        <label className="block font-medium text-white">Gross Total Income</label>
        <input type="number" name="grossIncome" value={formData.grossIncome} onChange={handleChange} className="w-full p-3 border rounded-lg shadow-sm" />
        
        <label className="block font-medium text-white">Income from Other Sources</label>
        <input type="number" name="otherIncome" value={formData.otherIncome} onChange={handleChange} className="w-full p-3 border rounded-lg shadow-sm" />
        
        <label className="block font-medium text-white">Employer PF</label>
        <input type="number" name="employerPf" value={formData.employerPf} onChange={handleChange} className="w-full p-3 border rounded-lg shadow-sm" />
        
        <label className="block font-medium text-white">Is Your Employer PF Separate?</label>
        <div className="flex gap-4">
          <label className="text-white"><input type="radio" name="isEmployerPfSeparate" value="true" checked={formData.isEmployerPfSeparate} onChange={handleChange} /> Yes</label>
          <label className="text-white"><input type="radio" name="isEmployerPfSeparate" value="false" checked={!formData.isEmployerPfSeparate} onChange={handleChange} /> No</label>
        </div>

        <label className="block font-medium text-white">Employee PF</label>
        <input type="number" name="employeePf" value={formData.employeePf} onChange={handleChange} className="w-full p-3 border rounded-lg shadow-sm" />
        
        <label className="block font-medium text-white">Do you want your employer to contribute to NPS?</label>
        <div className="flex gap-4">
          <label className="text-white"><input type="radio" name="employerNps" value="true" checked={formData.employerNps} onChange={handleChange} /> Yes</label>
          <label className="text-white"><input type="radio" name="employerNps" value="false" checked={!formData.employerNps} onChange={handleChange} /> No</label>
        </div>
        
        {formData.employerNps && (
          <>
            <label className="block font-medium text-white">Basic Income</label>
            <input type="number" name="basicIncome" value={formData.basicIncome} onChange={handleChange} className="w-full p-3 border rounded-lg shadow-sm" />
            
            <label className="block font-medium text-white">Percentage of Contribution (0-14%)</label>
            <input type="range" name="contributionPercentage" min="0" max="14" value={formData.contributionPercentage} onChange={handleChange} className="w-full" />
          </>
        )}

        <label className="block font-medium text-white">Do you want to opt for the old regime?</label>
        <div className="flex gap-4">
          <label className="text-white"><input type="radio" name="oldRegime" value="true" checked={formData.oldRegime} onChange={handleChange} /> Yes</label>
          <label className="text-white"><input type="radio" name="oldRegime" value="false" checked={!formData.oldRegime} onChange={handleChange} /> No</label>
        </div>

        {formData.oldRegime && (
          <>
            <label className="block font-medium text-white">Employee Contribution for NPS (0-50k)</label>
            <input type="number" name="employeeNps" value={formData.employeeNps} onChange={handleChange} className="w-full p-3 border rounded-lg shadow-sm" />

            <label className="block font-medium text-white">Interest from Saving Account</label>
            <input type="number" name="savingsInterest" value={formData.savingsInterest} onChange={handleChange} onBlur={handleSavingsBlur} className="w-full p-3 border rounded-lg shadow-sm" />
            {showSavingsMessage && formData.savingsInterest !== "" && (
              <p className="text-sm text-gray-300">
                {formData.savingsInterest <= 10000
                  ? `₹${formData.savingsInterest} is exempted under Section 12 TTA`
                  : `₹10000 is exempted under Section 12 TTA and ₹${formData.savingsInterest - 10000} is taxable as per slab`}
              </p>
            )}
          </>
        )}

        <button type="submit" className="w-full bg-white text-blue-600 p-3 rounded-lg shadow-md hover:bg-gray-200">Calculate Tax</button>
      </form>
    </div>
  );
}
