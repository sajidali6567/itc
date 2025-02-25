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
  
  const handleChange = (e) => {
    const { name, value, type, checked } = e.target;
    setFormData({
      ...formData,
      [name]: type === "checkbox" ? checked : value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const response = await fetch("http://localhost:8080/calculate-tax", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(formData),
    });
    const data = await response.json();
    console.log("Tax Calculation Result:", data);
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
        <input type="checkbox" name="isEmployerPfSeparate" checked={formData.isEmployerPfSeparate} onChange={handleChange} className="ml-2" />
        
        <label className="block font-medium text-white">Employee PF</label>
        <input type="number" name="employeePf" value={formData.employeePf} onChange={handleChange} className="w-full p-3 border rounded-lg shadow-sm" />
        
        <label className="block font-medium text-white">Do you want your employer to contribute to NPS?</label>
        <div className="flex space-x-4">
          <label><input type="radio" name="employerNps" value={false} checked={!formData.employerNps} onChange={handleChange} /> No</label>
          <label><input type="radio" name="employerNps" value={true} checked={formData.employerNps} onChange={handleChange} /> Yes</label>
        </div>
        {formData.employerNps && (
          <div>
            <label className="block font-medium text-white">Basic Income</label>
            <input type="number" name="basicIncome" value={formData.basicIncome} onChange={handleChange} className="w-full p-3 border rounded-lg shadow-sm" />
            <label className="block font-medium text-white">Percentage of Contribution (0-14%)</label>
            <input type="range" name="contributionPercentage" min="0" max="14" value={formData.contributionPercentage} onChange={handleChange} className="w-full" />
          </div>
        )}
        
        <label className="block font-medium text-white">Do you want to opt for the old regime?</label>
        <div className="flex space-x-4">
          <label><input type="radio" name="oldRegime" value={false} checked={!formData.oldRegime} onChange={handleChange} /> No</label>
          <label><input type="radio" name="oldRegime" value={true} checked={formData.oldRegime} onChange={handleChange} /> Yes</label>
        </div>
        {formData.oldRegime && (
          <div>
            <label className="block font-medium text-white">Employee Contribution for NPS (0-50k)</label>
            <input type="number" name="employeeNps" value={formData.employeeNps} onChange={handleChange} className="w-full p-3 border rounded-lg shadow-sm" />
            
            <label className="block font-medium text-white">Interest from Savings Account</label>
            <input type="number" name="savingsInterest" value={formData.savingsInterest} onChange={handleChange} className="w-full p-3 border rounded-lg shadow-sm" />
            <small className="text-white">10k exempted under Section 12 TTA; rest are taxable</small>
            
            <label className="block font-medium text-white">Income from Capital Gain</label>
            <div className="flex space-x-4">
              <label><input type="radio" name="capitalGain" value={false} checked={!formData.capitalGain} onChange={handleChange} /> No</label>
              <label><input type="radio" name="capitalGain" value={true} checked={formData.capitalGain} onChange={handleChange} /> Yes</label>
            </div>
            {formData.capitalGain && (
              <div>
                <label className="block font-medium text-white">Short-term Capital Gain (Taxable at 20%)</label>
                <input type="number" name="shortTermGain" value={formData.shortTermGain} onChange={handleChange} className="w-full p-3 border rounded-lg shadow-sm" />
                
                <label className="block font-medium text-white">Long-term Capital Gain (Exempted up to 1.25L, taxed at 12.5%)</label>
                <input type="number" name="longTermGain" value={formData.longTermGain} onChange={handleChange} className="w-full p-3 border rounded-lg shadow-sm" />
                
                <label className="block font-medium text-white">Dividend Earned (Taxable as per slab)</label>
                <input type="number" name="dividendEarned" value={formData.dividendEarned} onChange={handleChange} className="w-full p-3 border rounded-lg shadow-sm" />
              </div>
            )}
          </div>
        )}
        <button type="submit" className="w-full bg-white text-blue-600 p-3 rounded-lg shadow-md hover:bg-gray-200">Calculate Tax</button>
      </form>
    </div>
  );
}
