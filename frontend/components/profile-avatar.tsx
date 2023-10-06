"use client";
import {
  Avatar,
  Box,
  Button,
  Center,
  Flex,
  Modal,
  Text,
  TextInput,
  Textarea,
} from "@mantine/core";
import { useState } from "react";

export default function ProfileAvatar() {
  const [modalOpen, setModalOpen] = useState(false);

  return (
    <>
      <Modal
        opened={modalOpen}
        centered
        onClose={() => setModalOpen(false)}
        size="lg"
        title={<Text fw="bold">Profile</Text>}
        styles={{
          body: {
            padding: 0,
            borderTop: "1px solid #989FCE66",
          },
        }}
      >
        <form className="flex flex-col gap-5">
          <Box className="py-10 px-4 flex flex-col gap-5">
            <Center>
              <Avatar size="xl" />
            </Center>

            <TextInput label="First name" size="lg" />
            <TextInput label="Last name" size="lg" />
            <TextInput label="Address" size="lg" />
            <Textarea label="Bio" size="lg" />
          </Box>

          <Flex justify="end" className="py-2 border-t px-4">
            <Button size="md">Save</Button>
          </Flex>
        </form>
      </Modal>
      <Box pl={8} className="border-l">
        <Avatar
          size="md"
          onClick={() => setModalOpen(true)}
          className="cursor-pointer"
        />
      </Box>
    </>
  );
}
