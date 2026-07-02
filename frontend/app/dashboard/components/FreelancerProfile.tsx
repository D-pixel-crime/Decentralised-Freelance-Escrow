"use client";

import React, { useEffect, useState } from "react";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import axios from "axios";
import { Loader2, Save, User, Link as LinkIcon, BookOpen, Briefcase, Code, Github } from "lucide-react";

interface FreelancerProfileData {
  bio: string;
  resumeLink: string;
  experience: string;
  education: string;
  techStack: string[];
  githubLink: string;
  leetcodeLink: string;
  codeforcesLink: string;
}

export default function FreelancerProfile() {
  const queryClient = useQueryClient();
  const [success, setSuccess] = useState(false);

  const { data: profile, isLoading, isError } = useQuery<FreelancerProfileData>({
    queryKey: ["freelancerProfile"],
    queryFn: async () => {
      const res = await axios.get(
        `${process.env.NEXT_PUBLIC_API_URL}/api/get/profile`,
        { withCredentials: true }
      );
      return res.data;
    },
  });

  const [formData, setFormData] = useState<FreelancerProfileData>({
    bio: "",
    resumeLink: "",
    experience: "",
    education: "",
    techStack: [],
    githubLink: "",
    leetcodeLink: "",
    codeforcesLink: "",
  });

  const [techStackInput, setTechStackInput] = useState("");

  useEffect(() => {
    if (profile) {
      setFormData(profile);
      setTechStackInput(profile.techStack?.join(", ") || "");
    }
  }, [profile]);

  const mutation = useMutation({
    mutationFn: async (data: FreelancerProfileData) => {
      await axios.post(
        `${process.env.NEXT_PUBLIC_API_URL}/api/post/profile/update`,
        data,
        { withCredentials: true }
      );
    },
    onSuccess: () => {
      setSuccess(true);
      queryClient.invalidateQueries({ queryKey: ["freelancerProfile"] });
      setTimeout(() => setSuccess(false), 3000);
    },
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleTechStackChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setTechStackInput(e.target.value);
    setFormData((prev) => ({
      ...prev,
      techStack: e.target.value.split(",").map((s) => s.trim()).filter(Boolean),
    }));
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    mutation.mutate(formData);
  };

  if (isLoading) {
    return (
      <div className="flex flex-col items-center justify-center gap-3 py-32">
        <Loader2 className="h-8 w-8 animate-spin" style={{ color: `rgba(var(--vault-accent), 0.6)` }} />
        <span className="text-sm text-slate-500">Loading your profile...</span>
      </div>
    );
  }

  if (isError) {
    return (
      <div className="flex flex-col items-center justify-center gap-3 py-32">
        <div className="flex h-14 w-14 items-center justify-center rounded-full bg-red-500/10 border border-red-500/20">
          <span className="text-2xl">⚠</span>
        </div>
        <p className="text-sm text-slate-400">
          Failed to load profile. Please try again.
        </p>
      </div>
    );
  }

  return (
    <div className="mx-auto max-w-3xl">
      <div
        className="rounded-2xl p-8 backdrop-blur-sm shadow-2xl transition-all duration-300"
        style={{
          border: `1px solid rgba(var(--vault-accent), 0.15)`,
          background: `linear-gradient(145deg, rgba(15, 23, 42, 0.8), rgba(10, 14, 26, 0.9))`,
        }}
      >
        <div className="mb-8 flex items-center gap-4 border-b border-slate-800/60 pb-6">
          <div
            className="flex h-12 w-12 items-center justify-center rounded-xl"
            style={{
              background: `rgba(var(--vault-accent), 0.12)`,
              border: `1px solid rgba(var(--vault-accent), 0.25)`,
            }}
          >
            <User className="h-6 w-6" style={{ color: `rgba(var(--vault-accent), 1)` }} />
          </div>
          <div>
            <h2 className="text-xl font-bold text-slate-100">My Profile</h2>
            <p className="text-sm text-slate-400 mt-1">Manage your public freelancer information</p>
          </div>
        </div>

        <form onSubmit={handleSubmit} className="space-y-6">
          {/* Bio */}
          <div>
            <label className="mb-2 flex items-center gap-2 text-sm font-medium text-slate-300">
              <User className="h-4 w-4 text-slate-500" /> Bio
            </label>
            <textarea
              name="bio"
              value={formData.bio}
              onChange={handleChange}
              rows={3}
              placeholder="A short introduction about yourself..."
              className="w-full resize-none rounded-xl border border-slate-700/60 bg-slate-900/60 px-4 py-3 text-sm text-slate-200 placeholder-slate-600 outline-none transition-colors focus:border-[rgba(var(--vault-accent),0.5)] focus:ring-1 focus:ring-[rgba(var(--vault-accent),0.25)]"
            />
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
             {/* Tech Stack */}
             <div>
              <label className="mb-2 flex items-center gap-2 text-sm font-medium text-slate-300">
                <Code className="h-4 w-4 text-slate-500" /> Tech Stack
              </label>
              <input
                type="text"
                value={techStackInput}
                onChange={handleTechStackChange}
                placeholder="React, Solidity, Go..."
                className="w-full rounded-xl border border-slate-700/60 bg-slate-900/60 px-4 py-3 text-sm text-slate-200 placeholder-slate-600 outline-none transition-colors focus:border-[rgba(var(--vault-accent),0.5)] focus:ring-1 focus:ring-[rgba(var(--vault-accent),0.25)]"
              />
              <p className="mt-1.5 text-[11px] text-slate-500">Comma separated values</p>
            </div>
            
            {/* Resume Link */}
            <div>
              <label className="mb-2 flex items-center gap-2 text-sm font-medium text-slate-300">
                <LinkIcon className="h-4 w-4 text-slate-500" /> Resume Link
              </label>
              <input
                type="url"
                name="resumeLink"
                value={formData.resumeLink}
                onChange={handleChange}
                placeholder="https://..."
                className="w-full rounded-xl border border-slate-700/60 bg-slate-900/60 px-4 py-3 text-sm text-slate-200 placeholder-slate-600 outline-none transition-colors focus:border-[rgba(var(--vault-accent),0.5)] focus:ring-1 focus:ring-[rgba(var(--vault-accent),0.25)]"
              />
            </div>
          </div>

          {/* Experience */}
          <div>
            <label className="mb-2 flex items-center gap-2 text-sm font-medium text-slate-300">
              <Briefcase className="h-4 w-4 text-slate-500" /> Experience
            </label>
            <textarea
              name="experience"
              value={formData.experience}
              onChange={handleChange}
              rows={4}
              placeholder="Detail your relevant work experience..."
              className="w-full resize-none rounded-xl border border-slate-700/60 bg-slate-900/60 px-4 py-3 text-sm text-slate-200 placeholder-slate-600 outline-none transition-colors focus:border-[rgba(var(--vault-accent),0.5)] focus:ring-1 focus:ring-[rgba(var(--vault-accent),0.25)]"
            />
          </div>

          {/* Education */}
          <div>
            <label className="mb-2 flex items-center gap-2 text-sm font-medium text-slate-300">
              <BookOpen className="h-4 w-4 text-slate-500" /> Education
            </label>
            <input
              type="text"
              name="education"
              value={formData.education}
              onChange={handleChange}
              placeholder="University / Degree / Certifications..."
              className="w-full rounded-xl border border-slate-700/60 bg-slate-900/60 px-4 py-3 text-sm text-slate-200 placeholder-slate-600 outline-none transition-colors focus:border-[rgba(var(--vault-accent),0.5)] focus:ring-1 focus:ring-[rgba(var(--vault-accent),0.25)]"
            />
          </div>

          <div className="border-t border-slate-800/60 pt-6 mt-6">
            <h3 className="text-sm font-semibold text-slate-300 mb-4 flex items-center gap-2">
               External Profiles
            </h3>
            <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
              {/* Github */}
              <div>
                <label className="mb-1.5 flex items-center gap-1.5 text-xs font-medium text-slate-400">
                  <Github className="h-3.5 w-3.5" /> GitHub
                </label>
                <input
                  type="url"
                  name="githubLink"
                  value={formData.githubLink}
                  onChange={handleChange}
                  placeholder="https://github.com/..."
                  className="w-full rounded-xl border border-slate-700/60 bg-slate-900/60 px-3 py-2 text-sm text-slate-200 placeholder-slate-600 outline-none transition-colors focus:border-[rgba(var(--vault-accent),0.5)] focus:ring-1 focus:ring-[rgba(var(--vault-accent),0.25)]"
                />
              </div>

              {/* LeetCode */}
              <div>
                <label className="mb-1.5 flex items-center gap-1.5 text-xs font-medium text-slate-400">
                  <Code className="h-3.5 w-3.5" /> LeetCode
                </label>
                <input
                  type="url"
                  name="leetcodeLink"
                  value={formData.leetcodeLink}
                  onChange={handleChange}
                  placeholder="https://leetcode.com/..."
                  className="w-full rounded-xl border border-slate-700/60 bg-slate-900/60 px-3 py-2 text-sm text-slate-200 placeholder-slate-600 outline-none transition-colors focus:border-[rgba(var(--vault-accent),0.5)] focus:ring-1 focus:ring-[rgba(var(--vault-accent),0.25)]"
                />
              </div>

              {/* Codeforces */}
              <div>
                <label className="mb-1.5 flex items-center gap-1.5 text-xs font-medium text-slate-400">
                  <Code className="h-3.5 w-3.5" /> Codeforces
                </label>
                <input
                  type="url"
                  name="codeforcesLink"
                  value={formData.codeforcesLink}
                  onChange={handleChange}
                  placeholder="https://codeforces.com/..."
                  className="w-full rounded-xl border border-slate-700/60 bg-slate-900/60 px-3 py-2 text-sm text-slate-200 placeholder-slate-600 outline-none transition-colors focus:border-[rgba(var(--vault-accent),0.5)] focus:ring-1 focus:ring-[rgba(var(--vault-accent),0.25)]"
                />
              </div>
            </div>
          </div>

          {mutation.isError && (
             <div className="rounded-xl border border-red-500/20 bg-red-500/10 px-4 py-3">
               <p className="text-sm text-red-400">
                 Failed to save profile. Please try again.
               </p>
             </div>
          )}

          {success && (
            <div className="flex items-center gap-2 rounded-xl border border-emerald-500/20 bg-emerald-500/10 px-4 py-3">
              <span className="h-2 w-2 rounded-full bg-emerald-400" />
              <span className="text-sm font-medium text-emerald-400">
                Profile updated successfully!
              </span>
            </div>
          )}

          <div className="pt-4 flex justify-end">
            <button
              type="submit"
              disabled={mutation.isPending}
              className="flex items-center gap-2 rounded-xl px-6 py-3 text-sm font-semibold text-white transition-all duration-200 disabled:cursor-not-allowed disabled:opacity-50"
              style={{
                background: `rgba(var(--vault-accent), 0.85)`,
                boxShadow: `0 0 24px rgba(var(--vault-accent), 0.20)`,
              }}
              onMouseEnter={(e) => {
                if (!mutation.isPending) {
                  e.currentTarget.style.background = `rgba(var(--vault-accent), 1)`;
                  e.currentTarget.style.boxShadow = `0 0 32px rgba(var(--vault-accent), 0.35)`;
                  e.currentTarget.style.transform = "translateY(-1px)";
                }
              }}
              onMouseLeave={(e) => {
                e.currentTarget.style.background = `rgba(var(--vault-accent), 0.85)`;
                e.currentTarget.style.boxShadow = `0 0 24px rgba(var(--vault-accent), 0.20)`;
                e.currentTarget.style.transform = "translateY(0)";
              }}
            >
              {mutation.isPending ? (
                <>
                  <Loader2 className="h-4 w-4 animate-spin" />
                  Saving...
                </>
              ) : (
                <>
                  <Save className="h-4 w-4" />
                  Save Profile
                </>
              )}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}
