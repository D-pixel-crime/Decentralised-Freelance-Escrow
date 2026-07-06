import { NextResponse } from 'next/server';

export async function POST(request: Request) {
  try {
    const formData = await request.formData();
    const file = formData.get('file');

    if (!file) {
      return NextResponse.json(
        { error: 'No file provided' },
        { status: 400 }
      );
    }

    const pinataResponse = await fetch('https://api.pinata.cloud/pinning/pinFileToIPFS', {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${process.env.PINATA_JWT}`,
      },
      body: formData,
    });

    if (!pinataResponse.ok) {
      const errorText = await pinataResponse.text();
      return NextResponse.json(
        { error: 'Failed to upload to Pinata', details: errorText },
        { status: pinataResponse.status }
      );
    }

    const data = await pinataResponse.json();

    return NextResponse.json(
      { IpfsHash: data.IpfsHash },
      { status: 200 }
    );
  } catch (error) {
    console.error('Error uploading file to IPFS:', error);
    return NextResponse.json(
      { error: 'Internal server error' },
      { status: 500 }
    );
  }
}
