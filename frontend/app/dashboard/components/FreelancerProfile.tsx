"use client";

import React, { useEffect, useState } from "react";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import axios from "axios";
import { Loader2, Save, User, Link as LinkIcon, BookOpen, Briefcase, Code, Github } from "lucide-react";
import { useToast } from "@/contexts/ToastContext";
import { extractErrorMsg } from "@/lib/utils";

interface FreelancerProfileData {
  bio: string;
  resumeLink: string;
  experience: string;
  education: string;
  techStack: string[];
  githubLink: string;
  leetcodeLink: string;
  codeforcesLink: string;
  documentCids: string[];
}

export default function FreelancerProfile() {
  const queryClient = useQueryClient();
  const toast = useToast();

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
    documentCids: [],
  });

  const [techStackInput, setTechStackInput] = useState("");
  const [isUploadingFile, setIsUploadingFile] = useState(false);

  useEffect(() => {
    if (profile) {
      setFormData({
        ...profile,
        documentCids: profile.documentCids || [],
      });
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
      toast.success("Profile updated successfully!");
      queryClient.invalidateQueries({ queryKey: ["freelancerProfile"] });
    },
    onError: () => {
      toast.error("Failed to save profile. Please try again.");
    }
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

  const handleFileUpload = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;

    if (formData.documentCids.length >= 3) return;

    setIsUploadingFile(true);
    try {
      const uploadFormData = new FormData();
      uploadFormData.append("file", file);

      const res = await axios.post("/api/ipfs/file", uploadFormData);

      if (res.data?.IpfsHash) {
        setFormData(prev => ({
          ...prev,
          documentCids: [...prev.documentCids, res.data.IpfsHash]
        }));
      }
    } catch (err) {
      toast.error("Failed to upload file: " + extractErrorMsg(err));
    } finally {
      setIsUploadingFile(false);
      e.target.value = '';
    }
  };

  const removeDocument = (indexToRemove: number) => {
    setFormData(prev => ({
      ...prev,
      documentCids: prev.documentCids.filter((_, index) => index !== indexToRemove)
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
    <div className="mx-auto max-w-4xl">
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

          {/* Supporting Documents */}
          <div className="border-t border-slate-800/60 pt-6 mt-6">
            <h3 className="text-sm font-semibold text-slate-300 mb-4 flex items-center gap-2">
              Supporting Documents (Max 3)
            </h3>

            <div className="space-y-4">
              <div className="flex items-center gap-4">
                <input
                  type="file"
                  id="document-upload"
                  onChange={handleFileUpload}
                  disabled={formData.documentCids.length >= 3 || isUploadingFile}
                  className="hidden"
                />
                <label
                  htmlFor="document-upload"
                  className={`flex items-center gap-2 rounded-xl px-4 py-2 text-sm font-medium transition-all duration-200 cursor-pointer border ${formData.documentCids.length >= 3 || isUploadingFile
                      ? "opacity-50 cursor-not-allowed border-slate-700 bg-slate-800 text-slate-500"
                      : "border-[rgba(var(--vault-accent),0.4)] bg-[rgba(var(--vault-accent),0.1)] text-[rgba(var(--vault-accent),1)] hover:bg-[rgba(var(--vault-accent),0.2)]"
                    }`}
                >
                  {isUploadingFile ? (
                    <>
                      <Loader2 className="h-4 w-4 animate-spin" />
                      Uploading...
                    </>
                  ) : (
                    <>
                      <Save className="h-4 w-4" />
                      Upload Document
                    </>
                  )}
                </label>
                <span className="text-xs text-slate-500">
                  {formData.documentCids.length} / 3 uploaded
                </span>
              </div>

              {formData.documentCids.length > 0 && (
                <div className="grid gap-2">
                  {formData.documentCids.map((cid, index) => (
                    <div key={cid + index} className="flex items-center justify-between p-3 rounded-lg bg-slate-900/60 border border-slate-700/50">
                      <a
                        href={`https://gateway.pinata.cloud/ipfs/${cid}`}
                        target="_blank"
                        rel="noopener noreferrer"
                        className="flex items-center gap-2 text-sm text-slate-300 hover:text-[rgba(var(--vault-accent),1)] transition-colors truncate max-w-[80%]"
                      >
                        <LinkIcon className="h-4 w-4 shrink-0" />
                        <span className="truncate">{cid}</span>
                      </a>
                      <button
                        type="button"
                        onClick={() => removeDocument(index)}
                        className="text-xs text-red-400 hover:text-red-300 px-2 py-1 rounded bg-red-500/10 hover:bg-red-500/20 transition-colors"
                      >
                        Remove
                      </button>
                    </div>
                  ))}
                </div>
              )}
            </div>
          </div>
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
