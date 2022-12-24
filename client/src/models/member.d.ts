export interface IMember {
  connectionId: string;
  nickname: string;
}

export interface IMemberCardProps {
  member: IMember;
}

export interface IConnectedMembers {
  members: IMember[]
}