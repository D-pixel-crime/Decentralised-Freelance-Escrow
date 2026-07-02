import React, { useState, useEffect } from "react";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import axios from "axios";
import type { Job } from "@/types/job";
import {
  X, Loader2, User, Code, Link as LinkIcon, BookOpen, Briefcase, Github, Rocket
} from "lucide-react";
import { useAccount } from "wagmi";

function truncateAddress(addr: string) {
  if (!addr || addr.length < 12) return addr || "—";
  return `${addr.slice(0, 6)}…${addr.slice(-4)}`;
}

export default function ApplicantReviewModal({
  isOpen,
  onClose,
  job,
}: {
  isOpen: boolean;
  onClose: () => void;
  job: Job | null;
}) {
  const [selectedApplicant, setSelectedApplicant] = useState<string | null>(null);
  const [sliderValue, setSliderValue] = useState<number>(0);
  const queryClient = useQueryClient();
  const { address: clientAddress } = useAccount();

  useEffect(() => {
    if (isOpen && job && job.payMin !== undefined) {
      setSliderValue(job.payMin);
    }
  }, [isOpen, job]);

  // Fetch selected applicant's profile
  const { data: profile, isLoading: isProfileLoading, isError: isProfileError } = useQuery({
    queryKey: ["freelancerProfile", selectedApplicant],
    queryFn: async () => {
      const res = await axios.get(
        `${process.env.NEXT_PUBLIC_API_URL}/api/get/profile/${selectedApplicant}`,
        { withCredentials: true }
      );
      return res.data;
    },
    enabled: !!selectedApplicant,
    retry: false,
  });

  // Allocation mutation
  const allocateMutation = useMutation({
    mutationFn: async () => {
      if (!selectedApplicant || !job) throw new Error("Missing data");
      await axios.post(
        `${process.env.NEXT_PUBLIC_API_URL}/api/post/job/allocate`,
        {
          freelancerEthAccount: selectedApplicant,
          jobId: job.id,
          chainId: 31337, // Assuming Anvil/local for now
        },
        { withCredentials: true }
      );
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["myJobs"] });
      onClose(); // Close modal upon successful allocation
    },
  });

  if (!isOpen || !job) return null;

  return (
    <div
      className="fixed inset-0 z-50 flex items-center justify-center p-4"
      style={{ backdropFilter: "blur(8px)" }}
    >
      {/* Backdrop */}
      <div
        className="absolute inset-0 bg-black/60"
        onClick={onClose}
      />

      {/* Modal Container */}
      <div
        className="relative w-full max-w-5xl h-[85vh] rounded-2xl flex flex-col md:flex-row shadow-2xl overflow-hidden"
        style={{
          background: `linear-gradient(145deg, rgba(15, 23, 42, 0.95), rgba(10, 14, 26, 0.98))`,
          border: `1px solid rgba(var(--vault-accent), 0.20)`,
          boxShadow: `0 0 60px rgba(var(--vault-accent), 0.10), 0 25px 50px rgba(0,0,0,0.5)`,
        }}
      >
        {/* Close button */}
        <button
          onClick={onClose}
          className="absolute top-4 right-4 z-10 rounded-lg p-1.5 text-slate-500 transition-colors hover:bg-slate-800 hover:text-slate-300"
        >
          <X className="h-5 w-5" />
        </button>

        {/* Sidebar: Applicants List */}
        <div className="w-full md:w-1/3 border-b md:border-b-0 md:border-r border-slate-800/60 flex flex-col h-1/3 md:h-full bg-slate-900/40">
          <div className="p-5 border-b border-slate-800/60">
            <h2 className="text-lg font-bold tracking-tight text-slate-100">
              Applicants
            </h2>
            <p className="text-xs text-slate-500 mt-1">Select an applicant to review their profile</p>
          </div>
          <div className="flex-1 overflow-y-auto p-4 space-y-2">
            {job.applicants?.map((applicantAddr) => (
              <button
                key={applicantAddr}
                onClick={() => setSelectedApplicant(applicantAddr)}
                className={`w-full text-left p-3 rounded-xl border transition-all duration-200 flex items-center gap-3 ${
                  selectedApplicant === applicantAddr
                    ? "bg-[rgba(var(--vault-accent),0.1)] border-[rgba(var(--vault-accent),0.5)]"
                    : "bg-slate-800/30 border-slate-700/30 hover:bg-slate-800/50 hover:border-slate-700/50"
                }`}
              >
                <div className={`w-8 h-8 rounded-full flex items-center justify-center ${selectedApplicant === applicantAddr ? 'bg-[rgba(var(--vault-accent),0.2)]' : 'bg-slate-700'}`}>
                   <User className={`h-4 w-4 ${selectedApplicant === applicantAddr ? 'text-[rgba(var(--vault-accent),1)]' : 'text-slate-400'}`} />
                </div>
                <div>
                  <div className={`font-mono text-sm ${selectedApplicant === applicantAddr ? 'text-[rgba(var(--vault-accent),1)] font-semibold' : 'text-slate-300'}`}>
                    {truncateAddress(applicantAddr)}
                  </div>
                </div>
              </button>
            ))}
            {(!job.applicants || job.applicants.length === 0) && (
              <div className="text-center py-8 text-sm text-slate-500">
                No applicants yet.
              </div>
            )}
          </div>
        </div>

        {/* Main Content: Profile Review */}
        <div className="w-full md:w-2/3 flex flex-col h-2/3 md:h-full relative">
           {selectedApplicant ? (
             isProfileLoading ? (
               <div className="flex-1 flex flex-col items-center justify-center gap-3">
                 <Loader2 className="h-8 w-8 animate-spin" style={{ color: `rgba(var(--vault-accent), 0.6)` }} />
                 <span className="text-sm text-slate-500">Loading profile...</span>
               </div>
             ) : isProfileError ? (
               <div className="flex-1 flex flex-col items-center justify-center gap-3">
                 <p className="text-sm text-slate-400">Failed to load profile or freelancer hasn't set one up.</p>
               </div>
             ) : (
               <>
                 {/* Profile Header */}
                 <div className="p-6 md:p-8 flex-1 overflow-y-auto">
                    <h2 className="text-2xl font-bold text-slate-100 mb-6">Freelancer Profile</h2>
                    
                    <div className="space-y-8">
                       {/* Bio */}
                       <div>
                         <h3 className="text-sm font-semibold text-slate-300 mb-2 flex items-center gap-2">
                           <User className="h-4 w-4 text-slate-500" /> Bio
                         </h3>
                         <p className="text-sm text-slate-400 leading-relaxed whitespace-pre-wrap">
                           {profile?.bio || "No bio provided."}
                         </p>
                       </div>

                       {/* Experience & Education */}
                       <div className="grid md:grid-cols-2 gap-6">
                         <div>
                           <h3 className="text-sm font-semibold text-slate-300 mb-2 flex items-center gap-2">
                             <Briefcase className="h-4 w-4 text-slate-500" /> Experience
                           </h3>
                           <p className="text-sm text-slate-400 leading-relaxed whitespace-pre-wrap">
                             {profile?.experience || "Not specified."}
                           </p>
                         </div>
                         <div>
                           <h3 className="text-sm font-semibold text-slate-300 mb-2 flex items-center gap-2">
                             <BookOpen className="h-4 w-4 text-slate-500" /> Education
                           </h3>
                           <p className="text-sm text-slate-400 leading-relaxed whitespace-pre-wrap">
                             {profile?.education || "Not specified."}
                           </p>
                         </div>
                       </div>

                       {/* Tech Stack */}
                       <div>
                         <h3 className="text-sm font-semibold text-slate-300 mb-3 flex items-center gap-2">
                           <Code className="h-4 w-4 text-slate-500" /> Tech Stack
                         </h3>
                         <div className="flex flex-wrap gap-2">
                            {profile?.techStack && profile.techStack.length > 0 ? (
                              profile.techStack.map((tech: string, i: number) => (
                                <span key={i} className="inline-flex items-center rounded-lg bg-slate-800/80 border border-slate-700 px-2.5 py-1 text-xs font-medium text-slate-300">
                                  {tech}
                                </span>
                              ))
                            ) : (
                              <span className="text-sm text-slate-500">Not specified.</span>
                            )}
                         </div>
                       </div>

                       {/* Links */}
                       <div>
                         <h3 className="text-sm font-semibold text-slate-300 mb-3 flex items-center gap-2">
                           <LinkIcon className="h-4 w-4 text-slate-500" /> Links
                         </h3>
                         <div className="flex flex-wrap gap-4">
                           {profile?.githubLink && (
                             <a href={profile.githubLink} target="_blank" rel="noreferrer" className="flex items-center gap-1.5 text-sm text-[rgba(var(--vault-accent),0.9)] hover:underline">
                               <Github className="h-4 w-4" /> GitHub
                             </a>
                           )}
                           {profile?.resumeLink && (
                             <a href={profile.resumeLink} target="_blank" rel="noreferrer" className="flex items-center gap-1.5 text-sm text-[rgba(var(--vault-accent),0.9)] hover:underline">
                               <LinkIcon className="h-4 w-4" /> Resume
                             </a>
                           )}
                           {profile?.leetcodeLink && (
                             <a href={profile.leetcodeLink} target="_blank" rel="noreferrer" className="flex items-center gap-1.5 text-sm text-[rgba(var(--vault-accent),0.9)] hover:underline">
                               <Code className="h-4 w-4" /> LeetCode
                             </a>
                           )}
                           {profile?.codeforcesLink && (
                             <a href={profile.codeforcesLink} target="_blank" rel="noreferrer" className="flex items-center gap-1.5 text-sm text-[rgba(var(--vault-accent),0.9)] hover:underline">
                               <Code className="h-4 w-4" /> Codeforces
                             </a>
                           )}
                           {!profile?.githubLink && !profile?.resumeLink && !profile?.leetcodeLink && !profile?.codeforcesLink && (
                             <span className="text-sm text-slate-500">No links provided.</span>
                           )}
                         </div>
                       </div>
                    </div>
                 </div>

                 {/* Deployment Footer */}
                 <div className="p-6 border-t border-slate-800/60 bg-slate-900/40">
                   <div className="mb-4">
                     <label className="block text-xs font-semibold text-slate-400 mb-2 uppercase tracking-wider">
                       Allocation Setup
                     </label>
                     {/* UX Element: Stake amount slider/input (purely cosmetic per prompt constraint) */}
                     <div className="flex items-center gap-4">
                        <div className="flex-1">
                          <input 
                            type="range" 
                            className="w-full accent-[rgba(var(--vault-accent),1)] cursor-pointer disabled:cursor-not-allowed disabled:opacity-50" 
                            min={job?.payMin || 0}
                            max={job?.payMax || 0}
                            step="0.01"
                            value={sliderValue}
                            onChange={(e) => setSliderValue(Number(e.target.value))}
                            disabled={job?.payMin === job?.payMax}
                          />
                        </div>
                        <div className="text-xs text-slate-500 whitespace-nowrap">
                          {sliderValue} ETH (Stake set in next step)
                        </div>
                     </div>
                   </div>

                   <button
                    onClick={() => allocateMutation.mutate()}
                    disabled={allocateMutation.isPending || !clientAddress}
                    className="flex w-full items-center justify-center gap-2 rounded-xl px-4 py-3.5 text-sm font-semibold text-white transition-all duration-200 disabled:cursor-not-allowed disabled:opacity-50"
                    style={{
                      background: `rgba(var(--vault-accent), 0.85)`,
                      boxShadow: `0 0 24px rgba(var(--vault-accent), 0.20)`,
                    }}
                   >
                     {allocateMutation.isPending ? (
                       <>
                         <Loader2 className="h-5 w-5 animate-spin" />
                         Deploying Escrow Contract...
                       </>
                     ) : (
                       <>
                         <Rocket className="h-5 w-5" />
                         Select Freelancer & Deploy Contract
                       </>
                     )}
                   </button>
                   {allocateMutation.isError && (
                     <p className="mt-2 text-xs text-center text-red-400">
                       Failed to allocate job. Please ensure your wallet is connected and try again.
                     </p>
                   )}
                 </div>
               </>
             )
           ) : (
             <div className="flex-1 flex items-center justify-center text-slate-500 text-sm">
               Select an applicant from the list to review their profile.
             </div>
           )}
        </div>
      </div>
    </div>
  );
}
